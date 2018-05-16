package grains

import (
	"errors"
)

func Square(square int) (uint64, error) {
	if square < 1 || square > 64 {
		return 0, errors.New("square should be between 1 and 64")
	}
	return 1 << uint64(square-1), nil
}
func Total() uint64 {
	return ^uint64(0)
}