package strain

type Ints []int
type Lists [][]int
type Strings []string

func (is Ints) Keep(f func(int) bool) Ints {
	var keep Ints
	for _, i := range is {
		if f(i) {
			keep = append(keep, i)
		}
	}
	return keep
}

func (is Ints) Discard(f func(int) bool) Ints {
	var keep Ints
	for _, i := range is {
		if !f(i) {
			keep = append(keep, i)
		}
	}
	return keep
}

func (ls Lists) Keep(f func([]int) bool) Lists {
	var keep Lists
	for _, l := range ls {
		if f(l) {
			keep = append(keep, l)
		}
	}
	return keep
}

func (ss Strings) Keep(f func(string) bool) Strings {
	var keep Strings
	for _, s := range ss {
		if f(s) {
			keep = append(keep, s)
		}
	}
	return keep
}