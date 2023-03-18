package listops

type IntList []int

// Foldl, given a function, a list, and initial accumulator, folds (reduce) each item into the accumulator
// from the left using `function(accumulator, item)`
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, num := range s {
		initial = fn(initial, num)
	}
	return initial
}

// Foldr, given a function, a list, and an initial accumulator, folds (reduce) each item into the accumulator
// from the right using `function(item, accumulator)`
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

// Filter, given a predicate and a list, returns the list of all items for which `predicate(item)` is True
func (s IntList) Filter(fn func(int) bool) IntList {
	fl := []int{}
	for _, num := range s {
		if fn(num) {
			fl = append(fl, num)
		}
	}
	return fl
}

// Length, given a list, returns the total number of items within it
func (s IntList) Length() int {
	return len(s)
}

// Map, given a function and a list, returns the list of the results of applying
// `function(item)` on all items
func (s IntList) Map(fn func(int) int) IntList {
	ml := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		ml[i] = fn(s[i])
	}
	return ml
}

// Reverse, given a list, returns a list with all the original items, but in reversed order
func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Append, given two lists, adds all items in the second list to the end of the first list
func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

// Concat,  given a series of lists, combines all items in all lists into one flattened list
func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = s.Append(l)
	}
	return s
}
