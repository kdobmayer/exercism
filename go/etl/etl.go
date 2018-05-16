package etl

import "strings"

func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, letters := range input {
		for _, l := range letters {
			output[strings.ToLower(l)] = score
		}
	}
	return output
}