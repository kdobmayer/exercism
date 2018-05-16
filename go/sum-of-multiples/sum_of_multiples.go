package summultiples

func SumMultiples(limit int, divisor ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, d := range divisor {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}