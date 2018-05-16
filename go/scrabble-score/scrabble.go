package scrabble

import "unicode"

type Scrabble map[int][]rune

func (s Scrabble) Score(r rune) int {
	for score, letters := range s {
		for _, letter := range letters {
			if unicode.ToUpper(r) == letter {
				return score
			}
		}
	}
	return -1
}

var scrabble = Scrabble{
	1:  {'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	2:  {'D', 'G'},
	3:  {'B', 'C', 'M', 'P'},
	4:  {'F', 'H', 'V', 'W', 'Y'},
	5:  {'K'},
	8:  {'J', 'X'},
	10: {'Q', 'Z'},
}

func Score(word string) int {
	score := 0
	for _, r := range word {
		score += scrabble.Score(r)
	}
	return score
}