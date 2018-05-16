package transpose

import (
	"strings"
)

func Transpose(in []string) []string {
	out := make([]string, longest(in))
	for i := 0; i < len(out); i++ {
		for j := 0; j < len(in); j++ {
			out[i] += string(in[j][i])
		}
	}
	for i := len(out) - 1; i >= 0; i-- {
		if !strings.HasSuffix(out[i], " ") {
			break
		}
		out[i] = strings.TrimRight(out[i], " ")
	}
	return out
}

func longest(in []string) int {
	var max int
	for _, s := range in {
		if len(s) > max {
			max = len(s)
		}
	}
	for i, s := range in {
		in[i] += strings.Repeat(" ", max-len(s))
	}
	return max
}