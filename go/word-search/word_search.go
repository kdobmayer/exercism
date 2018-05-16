package wordsearch

import (
	"fmt"
)

func Solve(words, puzzle []string) (map[string][2][2]int, error) {
	coords := make(map[string][2][2]int, len(words))
	for _, word := range words {
		c, err := find(word, puzzle)
		if err != nil {
			return nil, fmt.Errorf("can't solve: %v", err)
		}
		coords[word] = c
	}
	return coords, nil
}

func find(word string, puzzle []string) ([2][2]int, error) {
	var coord [2][2]int
	rows, cols, length := len(puzzle), len(puzzle[0]), len(word)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i <= rows-length {
				switch text(puzzle, i, j, 1, 0, length) {
				case word: // top to bottom
					return [2][2]int{{j, i}, {j, i + length - 1}}, nil
				case reverse(word): // bottom to top
					return [2][2]int{{j, i + length - 1}, {j, i}}, nil
				}
				if j <= rows-length {
					switch text(puzzle, i, j, 1, 1, length) {
					case word: // top left to bottom right
						return [2][2]int{{j, i}, {j + length - 1, i + length - 1}}, nil
					case reverse(word): // bottom right to top left
						return [2][2]int{{j + length - 1, i + length - 1}, {j, i}}, nil
					}
				}
			}
			if j <= cols-length {
				switch text(puzzle, i, j, 0, 1, length) {
				case word: // left to right
					return [2][2]int{{j, i}, {j + length - 1, i}}, nil
				case reverse(word): // right to left
					return [2][2]int{{j + length - 1, i}, {j, i}}, nil
				}
				if i >= length-1 {
					switch text(puzzle, i, j, -1, 1, length) {
					case word: // bottom left to top right
						return [2][2]int{{j, i}, {j + length - 1, i - length + 1}}, nil
					case reverse(word): // top right to bottom left
						return [2][2]int{{j + length - 1, i - length + 1}, {j, i}}, nil
					}
				}
			}
		}
	}

	return coord, fmt.Errorf("word %q not found", word)
}

func text(puzzle []string, x, y, down, left, length int) string {
	var s string
	for i := 0; i < length; i++ {
		s += string(puzzle[x+i*down][y+i*left])
	}
	return s
}

func reverse(s string) string {
	r := make([]byte, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		r[i], r[j] = s[j], s[i]
	}
	return string(r)
}