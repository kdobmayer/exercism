package palindrome

import (
	"errors"
	"strconv"
)

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			num := i * j
			factor := [2]int{i, j}
			if isPalindrome(num) {
				if num < pmin.Product || pmin.Product == 0 {
					pmin.Product = num
					pmin.Factorizations = [][2]int{factor}
				} else if num == pmin.Product {
					pmin.Factorizations = append(pmin.Factorizations, factor)
				}
				if num > pmax.Product || pmax.Product == 0 {
					pmax.Product = num
					pmax.Factorizations = [][2]int{factor}
				} else if num == pmax.Product {
					pmax.Factorizations = append(pmax.Factorizations, factor)
				}
			}
		}
	}

	if pmin.Product == 0 || pmax.Product == 0 {
		return pmin, pmax, errors.New("no palindromes")
	}
	return pmin, pmax, nil
}

func isPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}