package sublist

import "reflect"

type Relation string

const (
	sublist   Relation = "sublist"
	superlist          = "superlist"
	equal              = "equal"
	unequal            = "unequal"
)

func Sublist(first, second []int) Relation {
	switch {

	case reflect.DeepEqual(first, second):
		return equal
	case isSublist(first, second):
		return sublist
	case isSublist(second, first):
		return superlist
	default:
		return unequal
	}
}

func isSublist(first, second []int) bool {
	for i := 0; i <= len(second)-len(first); i++ {
		if reflect.DeepEqual(second[i:i+len(first)], first) {
			return true
		}
	}
	return false
}