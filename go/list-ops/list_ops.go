package listops

type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, num := range s {
        initial = fn(initial, num)
    }
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i:= len(s)-1; i >= 0 ; i--{
        initial = fn(s[i], initial)
    }
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	fl := []int{}
    for _, num := range s {
        if fn(num) {
            fl  = append(fl, num)
        }
    }
	return fl
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	ml := make([]int, len(s))
    for i:= 0; i < len(s); i++ {
        ml[i] = fn(s[i])
    }
	return ml
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
        s = s.Append(l)
    }
	return s
}
