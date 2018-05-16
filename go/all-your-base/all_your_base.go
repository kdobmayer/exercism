package allyourbase

import "errors"

func ConvertToBase(from int, digits []int, to int) ([]int, error) {
	if from < 2 {
		return []int{}, errors.New("input base must be >= 2")
	} else if to < 2 {
		return []int{}, errors.New("output base must be >= 2")
	}

	decimal := 0
	for i, j := len(digits)-1, 1; i >= 0; i, j = i-1, j*from {
		if digits[i] < 0 || digits[i] >= from {
			return []int{}, errors.New("all digits must satisfy 0 <= d < input base")
		}
		decimal += digits[i] * j
	}

	var converted []int
	for i := 1; i <= decimal; i *= to {
		converted = append([]int{(decimal % (i * to)) / i}, converted...)
	}
	if len(converted) == 0 {
		converted = []int{0}
	}
	return converted, nil
}