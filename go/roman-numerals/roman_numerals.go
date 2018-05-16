package romannumerals

import (
	"errors"
	"strings"
)

// arabic to roman
var ator = map[int]string{
	1: "I", 5: "V",
	10: "X", 50: "L",
	100: "C", 500: "D",
	1000: "M",
}

func ToRomanNumeral(arabic int) (string, error) {
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("invalid arabic number")
	}
	var roman string
	for i := 1; i <= arabic; i *= 10 {
		switch digit := (arabic / i) % 10; digit {
		case 1, 2, 3:
			roman = strings.Repeat(ator[1*i], digit) + roman
		case 4, 9:
			roman = ator[1*i] + ator[(digit+1)*i] + roman
		case 5, 6, 7, 8:
			roman = ator[5*i] + strings.Repeat(ator[1*i], digit-5) + roman
		}
	}
	return roman, nil
}