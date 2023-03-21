package circular

import (
	"errors"
)

// Buffer represents a circular buffer, cyclic buffer or ring buffer, a data structure that
// uses a single, fixed-size buffer as if it is connected end-to-end
type Buffer struct {
	b    chan byte
	size int
}

// NewBuffer return a pointer to a Buffer with the given length
func NewBuffer(size int) *Buffer {
	buffer := Buffer{
		b:    make(chan byte, size),
		size: size,
	}
	return &buffer
}

// ReadByte removes the oldest value from the buffer and returns it.
// Returns an error if the buffer is empty
func (b *Buffer) ReadByte() (byte, error) {
	if len(b.b) == 0 {
		return 0, errors.New("empty buffer")
	}
	return <-b.b, nil
}

// WriteByte adds a value to the buffer, returns an error only if the buffer is full
func (b *Buffer) WriteByte(c byte) error {
	if len(b.b) == b.size {
		return errors.New("the buffer is full")
	}
	b.b <- c
	return nil
}

// Overwrite substitutes the oldest value in the buffer with
// the newly given one
func (b *Buffer) Overwrite(c byte) {
	err := b.WriteByte(c)
	if err != nil {
		<-b.b
		b.b <- c
	}
}

// Reset modifies the internal state of a Buffer to allocate  anew buffer for values
// of the previously determined length
func (b *Buffer) Reset() {
	b.b = make(chan byte, b.size)
}
