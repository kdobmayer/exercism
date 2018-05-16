package perfect

import "errors"

var ErrOnlyPositive = errors.New("input is not a positive integer")

type Classification int

const (
	ClassificationAbundant Classification = iota
	ClassificationDeficient
	ClassificationPerfect
)

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}
	switch sum := aliquotSum(n); {
	case sum < n:
		return ClassificationDeficient, nil
	case sum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationPerfect, nil
	}
}

func aliquotSum(n int64) int64 {
	var sum int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}