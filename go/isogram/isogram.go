package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, r := range strings.ToLower(word) {
		if !unicode.IsLetter(r) {
			continue
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}