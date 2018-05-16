package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var anagrams []string
	for _, c := range candidates {
		if isAnagram(subject, c) {
			anagrams = append(anagrams, c)
		}
	}
	return anagrams
}

func isAnagram(a, b string) bool {
	a, b = strings.ToLower(a), strings.ToLower(b)

	if a == b || len(a) != len(b) {
		return false
	}

	for _, r := range a {
		i := strings.IndexRune(b, r)
		if i == -1 {
			return false
		}
		b = b[:i] + b[i+1:]
	}
	return true
}