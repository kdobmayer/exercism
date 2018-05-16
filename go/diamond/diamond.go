package diamond

import (
	"errors"
	"fmt"
	"strings"
)

const A = int('A')

func Gen(b byte) (string, error) {
	if b < 'A' || b > 'Z' {
		return "", errors.New("invalid input")
	}
	target := int(b)
	var diamond []string
	for i := A; i <= target; i++ {
		diamond = append(diamond, row(i, target))
	}
	for i := target - 1; i >= A; i-- {
		diamond = append(diamond, row(i, target))
	}
	return strings.Join(diamond, "\n") + "\n", nil
}

func row(current, target int) string {
	var centrum string
	if current == 'A' {
		centrum = "A"
	} else {
		centrum = fmt.Sprintf("%c%s%[1]c", current, spaces(2*(current-A)-1))
	}

	return spaces(target-current) + centrum + spaces(target-current)
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}