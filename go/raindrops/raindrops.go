package raindrops

import (
	"strconv"
)

var rules = []struct {
	factor int
	output string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(n int) string {
	converted := ""
	for _, rule := range rules {
		if n%rule.factor == 0 {
			converted += rule.output
		}
	}
	if converted == "" {
		return strconv.Itoa(n)
	}
	return converted
}