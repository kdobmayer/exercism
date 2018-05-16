package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func Encode(pt string) string {
	pt = normalize(pt)

	var ct string
	rows := int(math.Ceil(math.Sqrt(float64(len(pt)))))
	for i := 0; i < rows; i++ {
		for j := 0; j < rows+1; j++ {
			if j*rows+i < len(pt) {
				ct += string(pt[j*rows+i])
			}
		}
		ct += " "
	}
	return strings.TrimSpace(ct)
}

func normalize(s string) string {
	var n string
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			n += string(unicode.ToLower(r))
		}
	}
	return n
}