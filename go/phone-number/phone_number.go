package phonenumber

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func Number(input string) (string, error) {
	var num string
	for _, r := range input {
		if unicode.IsDigit(r) {
			num += string(r)
		}
	}
	num = strings.TrimLeft(num, "1")
	if len(num) != 10 || num[3] < '2' {
		return "", errors.New("invalid input")
	}
	return num, nil
}

func Format(input string) (string, error) {
	num, err := Number(input)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[0:3], num[3:6], num[6:]), nil
}

func AreaCode(input string) (string, error) {
	num, err := Number(input)
	if err != nil {
		return "", err
	}
	return num[0:3], nil
}