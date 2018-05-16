package stringset

import (
	"fmt"
	"reflect"
)

type Set map[string]bool

func New() Set {
	return make(Set)
}

func NewFromSlice(a []string) Set {
	s := make(Set)
	for _, v := range a {
		s[v] = true
	}
	return s
}

func (s Set) String() string {
	str, sep := "{", ""
	for elem := range s {
		str += fmt.Sprintf("%s%q", sep, elem)
		sep = ", "
	}
	return str + "}"
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(e string) bool {
	return s[e]
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2[k] {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2[k] {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return reflect.DeepEqual(s1, s2)
}

func (s Set) Add(e string) {
	s[e] = true
}

func Intersection(s1, s2 Set) Set {
	inter := make(Set)
	for k := range s1 {
		if s2[k] {
			inter[k] = true
		}
	}
	return inter
}

func Difference(s1, s2 Set) Set {
	diff := make(Set)
	for k := range s1 {
		if !s2[k] {
			diff[k] = true
		}
	}
	return diff
}

func Union(s1, s2 Set) Set {
	uni := make(Set)
	for k := range s1 {
		uni[k] = true
	}
	for k := range s2 {
		uni[k] = true
	}
	return uni
}