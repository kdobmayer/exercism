package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Hammer distance is only defined for sequences of equal length.")
	}
	distance := 0
	for i, r := range a {
		if r != rune(b[i]) {
			distance++
		}
	}
	return distance, nil
}