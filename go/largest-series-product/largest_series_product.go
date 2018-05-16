package lsproduct

import (
	"errors"
	"fmt"
	"unicode"
)

func LargestSeriesProduct(input string, span int) (int, error) {
	digits, err := parse(input, span)
	if err != nil {
		return 0, err
	}

	var max int
	for i := 0; i <= len(digits)-span; i++ {
		product := 1
		for j := 0; j < span; j++ {
			product *= digits[i+j]
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}

func parse(input string, span int) ([]int, error) {
	var digits []int

	if len(input) < span || span < 0 {
		return digits, errors.New("Invalid input")
	}

	for _, r := range input {
		if !unicode.IsDigit(r) {
			return digits, fmt.Errorf("invalid digit: %c\n", r)
		}
		digits = append(digits, int(r-48))
	}
	return digits, nil
}