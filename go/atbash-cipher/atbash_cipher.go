package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var encrypted string
	for _, r := range strings.ToLower(s) {
		encrypted += cipher(r)
	}
	words := regexp.MustCompile(".{1,5}").FindAllString(encrypted, -1)
	return strings.Join(words, " ")
}

func cipher(r rune) string {
	if unicode.IsDigit(r) {
		return string(r)
	}
	if unicode.IsLetter(r) {
		return string('a' + 'z' - r)
	}
	return ""
}