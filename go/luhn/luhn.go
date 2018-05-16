package luhn

import (
	"strings"
	"unicode"
)

// Valid uses the Luhn algorithm to check identification numbers.
func Valid(input string) bool {
	// Spaces should be stripped before checking.
	input = strings.Replace(input, " ", "", -1)
	// Strings of length 1 or less are not valid.
	if len(input) <= 1 {
		return false
	}
	sum := 0
	// Calculate checksum starting from the right.
	for i, r := range input {
		// Non-digit characters are disallowed.
		if !unicode.IsDigit(r) {
			return false
		}
		digit := int(r - '0')
		if len(input)%2 == i%2 {
			digit = (2*digit)/10 + (2*digit)%10
		}
		sum += digit
	}
	// If the sum is evenly divisible by 10, then the number is valid.
	return sum%10 == 0
}