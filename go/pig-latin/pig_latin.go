package igpay

import (
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u", "yt", "xr"}
var consonants = []string{"ch", "qu", "squ", "thr", "th", "sch"}

func PigLatin(in string) string {
	words := strings.Split(in, " ")
	for i, word := range words {
		words[i] = pigLatin(word)
	}
	return strings.Join(words, " ")
}

func pigLatin(s string) string {
	if isFirstVowel(s) {
		return s + "ay"
	}
	for _, c := range consonants {
		if strings.HasPrefix(s, c) {
			return strings.TrimPrefix(s, c) + c + "ay"
		}
	}
	return s[1:] + s[:1] + "ay"
}

func isFirstVowel(s string) bool {
	for _, v := range vowels {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}