package railfence

import "strings"

func Encode(message string, rails int) string {
	lines := make([]string, rails)
	line, down := 0, true
	for _, r := range message {
		lines[line] += string(r)
		switch line {
		case 0:
			down = true
		case rails - 1:
			down = false
		}
		line = next(line, down)
	}
	return strings.Join(lines, "")
}

func Decode(message string, rails int) string {
	rows := fence(rails, message)
	letter := 0
	for i, row := range rows {
		for j, r := range row {
			if r == '?' {
				rows[i][j] = message[letter]
				letter++
			}
		}
	}
	var decoded []byte
	for i := 0; i < len(message); i++ {
		for _, row := range rows {
			if row[i] > 0 {
				decoded = append(decoded, row[i])
			}
		}
	}

	return string(decoded)
}

// fence creates a byte matrix where the question marks shows where the letters
// would be.
func fence(rails int, message string) [][]byte {
	rows := make([][]byte, rails)
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]byte, len(message))
	}
	line, down := 0, true
	for i := 0; i < len(message); i++ {
		rows[line][i] = '?'
		switch line {
		case 0:
			down = true
		case rails - 1:
			down = false
		}
		line = next(line, down)
	}
	return rows
}

// next provides the next line where append the next letter to. If down is true, then
// move one line downward, if false, then one line upward.
func next(line int, down bool) int {
	if down {
		return line + 1
	}
	return line - 1
}