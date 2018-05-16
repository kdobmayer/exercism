package sieve

func Sieve(limit int) []int {
	var primes []int
	composites := make(map[int]bool)
	for i := 2; i <= limit; i++ {
		if !composites[i] {
			primes = append(primes, i)
			for j := 2; j <= limit/i; j++ {
				composites[i*j] = true
			}
		}
	}
	return primes
}