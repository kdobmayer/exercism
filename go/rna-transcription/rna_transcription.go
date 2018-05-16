package strand

import "strings"

var comps = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	return strings.Map(func(r rune) rune {
		return comps[r]
	}, dna)
}