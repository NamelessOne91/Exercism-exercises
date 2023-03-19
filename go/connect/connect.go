package connect

import (
	"errors"
	"sync"
)

// Board represents the state of a bidimensional playing board for Hex
type Board []string

// Position represents the placement of a piece on the game's board
type Position struct {
	row, col int
	symbol   rune
}

// Move represents an available move on a bidimensional plane
type Move struct {
	R, C int
}

// Position contains all the available movements a piece can make on a bidimensional board
var moves = [...]Move{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

// Cache is a data structure holding up to the 2 previous positions of a piece
type Cache [2]Position

// hasWon return a boolean representing if the position is a winning one for that symbol
func (p Position) hasWon() bool {
	return p.symbol == 'X' && p.col == 0 || p.symbol == 'O' && p.row == 0
}

// isInside returns a boolean representing if a given position is inside the game's board
func (b Board) isInside(p Position) bool {
	return p.row < len(b) && p.row >= 0 && p.col < len(b[0]) && p.col >= 0
}

// isConnected returns a boolean representing if the given Position's symbol
// is the same one displayed on the board
func (b Board) isConnected(p Position) bool {
	return p.symbol == rune(b[p.row][p.col])
}

// validDiagonal returns a boolean representing if a 2 position are on valid diagonal
// for a bidimensional board
func (p Position) validDiagonal(newP Position) bool {
	return (p.col + p.row) == (newP.col + newP.row)
}

// isDiagonal returns a boolean to represent if the move is along the diagonal of a
// bidimensional board
func (m Move) isDiagonal() bool {
	delta := m.R + m.C
	return delta != 1 && delta != -1
}

// isLooping returns a boolean representing if a given position has already been visited
func (c Cache) isLooping(p Position) bool {
	// check against the 2 cached previous moves
	for _, prevPos := range c {
		// loop found
		if p.row == prevPos.row && p.col == prevPos.col {
			return true
		}
	}
	return false
}

// move updates the row and colums of a Position according to the given Move
// and returns a new Position and boolean to represent if it is a valid move
func (b Board) move(p Position, m Move) (Position, bool) {
	newP := Position{
		symbol: p.symbol,
		row:    p.row + m.R,
		col:    p.col + m.C,
	}

	if !b.isInside(newP) || (m.isDiagonal() && !p.validDiagonal(newP)) || !b.isConnected(newP) {
		return newP, false
	}
	return newP, true
}

// isWinner computes the state of a given Board and sends on the given channel
// the given Position's symbol if it is the winner
func notifyWinner(b Board, start Position, c Cache, resChan chan<- rune) {
	// base cases for winning
	if start.hasWon() {
		resChan <- start.symbol
		return
	}
	// check each available move
	for _, move := range moves {
		newP, ok := b.move(start, move)
		if ok && !c.isLooping(newP) {
			c[0], c[1] = start, c[0]
			notifyWinner(b, newP, c, resChan)
		}
	}
}

// ResultOf computes the state of a given board and returns a string
// representing the winning smbol or lack of it
func ResultOf(board Board) (string, error) {
	rows := len(board)
	if rows < 1 {
		return "", errors.New("no board state provided")
	}
	cols := len(board[0])

	cache := Cache{}
	var wg sync.WaitGroup
	var winChan = make(chan rune)

	go func() {
		for c, v := range board[rows-1] {
			if v == 'O' {
				wg.Add(1)
				go func(col int, symbol rune) {
					notifyWinner(board, Position{rows - 1, col, symbol}, cache, winChan)
					wg.Done()
				}(c, v)
			}
		}
		for r := 0; r < rows; r++ {
			v := rune(board[r][cols-1])
			if v == 'X' {
				wg.Add(1)
				go func(row int, symbol rune) {
					notifyWinner(board, Position{row, cols - 1, symbol}, cache, winChan)
					wg.Done()
				}(r, v)
			}
		}
		wg.Wait()
		close(winChan)
	}()

	for winner := range winChan {
		return string(winner), nil
	}
	return "", nil
}
