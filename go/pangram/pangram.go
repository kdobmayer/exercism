package pangram

import (
	"strings"
)

func IsPangram(s string) bool {
	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(strings.ToLower(s), string(i)) {
			return false
		}
	}
	return true
}