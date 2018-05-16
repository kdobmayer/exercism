package prime

import "math"

func Factors(n int64) []int64 {
	fs := []int64{}
	for i := int64(2); n != 1; i = NextPrime(i) {
		for n%i == 0 {
			fs = append(fs, i)
			n /= i
		}
	}
	return fs
}

func NextPrime(n int64) int64 {
	i := n + 1
	for !IsPrime(i) {
		i++
	}
	return i
}

func IsPrime(n int64) bool {
	for i := int64(2); i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}