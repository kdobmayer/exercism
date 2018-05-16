package say

import (
	"fmt"
	"math"
	"strings"
)

var letters = map[int]string{
	//0:    "zero",
	1:   "one",
	2:   "two",
	3:   "three",
	4:   "four",
	5:   "five",
	6:   "six",
	7:   "seven",
	8:   "eight",
	9:   "nine",
	10:  "ten",
	11:  "eleven",
	12:  "twelve",
	13:  "thirteen",
	14:  "fourteen",
	15:  "fifteen",
	16:  "sixteen",
	17:  "seventeen",
	18:  "eighteen",
	19:  "nineteen",
	20:  "twenty",
	30:  "thirty",
	40:  "forty",
	50:  "fifty",
	60:  "sixty",
	70:  "seventy",
	80:  "eighty",
	90:  "ninety",
	100: "one hundred",
	1e3: "thousand",
	1e6: "million",
	1e9: "billion",
}

func Say(n int64) (string, bool) {
	if n < 0 || n >= 1e12 {
		return "", false
	} else if n == 0 {
		return "zero", true
	}

	var s string
	for i := 9; i > 1; i -= 3 {
		chunk := int(n) % int(math.Pow10(i+3)) / int(math.Pow10(i))
		if chunk >= 1 {
			s += fmt.Sprintf("%s %s ", number(chunk), letters[int(math.Pow10(i))])
		}
	}
	s += number(int(n) % 1e3)
	return strings.TrimSpace(s), true
}

func number(n int) string {
	var s string

	hundreds := n / 100
	if hundreds >= 1 {
		s += fmt.Sprintf("%s hundred ", letters[hundreds])
	}

	decimal := n % 100
	if decimal < 20 {
		s += letters[decimal]
		return s
	}

	s += letters[(decimal/10)*10]
	if decimal%10 != 0 {
		s += "-" + letters[decimal%10]
	}
	return s
}