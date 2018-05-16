package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var acronym string
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	for _, w := range words {
		acronym += string(w[0])
	}
	return strings.ToUpper(acronym)
}