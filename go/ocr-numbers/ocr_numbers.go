package ocr

import (
	"strings"
)

func recognizeDigit(s string) string {
	switch s {
	case "_| ||_|":
		return "0"
	case "   |  |":
		return "1"
	case "_ _||_ ":
		return "2"
	case "_ _| _|":
		return "3"
	case " |_|  |":
		return "4"
	case "_|_  _|":
		return "5"
	case "_|_ |_|":
		return "6"
	case "_  |  |":
		return "7"
	case "_|_||_|":
		return "8"
	case "_|_| _|":
		return "9"
	default:
		return "?"
	}
}

func Recognize(s string) []string {
	var numbers []string
	lines := strings.Split(s, "\n")[1:]
	for i := 0; i < len(lines); i += 4 {
		var num string
		for j := 0; j < len(lines[0]); j += 3 {
			num += recognizeDigit(lines[i][j+1:j+2] + lines[i+1][j:j+3] + lines[i+2][j:j+3])
		}
		numbers = append(numbers, num)
	}
	return numbers
}