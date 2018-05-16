package prime

import (
	"math"
)

func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	p := 2 // first prime
	for i := 2; i <= n; i++ {
		p = NextPrime(p)
	}
	return p, true
}

func NextPrime(n int) int {
	i := n + 1
	for !IsPrime(i) {
		i++
	}
	return i
}

func IsPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}