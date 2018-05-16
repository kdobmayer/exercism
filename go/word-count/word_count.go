package wordcount

import (
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(input string) Frequency {
	freq := make(Frequency)
	words := strings.FieldsFunc(input, func(r rune) bool { return unicode.IsSpace(r) || r == ',' })
	for _, w := range words {
		w = strings.TrimFunc(w, func(r rune) bool { return !unicode.IsLetter(r) && !unicode.IsDigit(r) })
		freq[strings.ToLower(w)]++
	}
	return freq
}