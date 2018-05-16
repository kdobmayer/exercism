package minesweeper

import (
	"bytes"
	"errors"
	"regexp"
)

type Board [][]byte

func (b Board) String() string {
	return "\n" + string(bytes.Join(b, []byte{'\n'}))
}

func (b Board) Count() error {
	cols := len(bytes.Join(b, []byte{})) / len(b)
	for i, row := range b {
		if len(row) != cols {
			return errors.New("unaligned")
		}
		var pattern string
		if i == 0 || i == len(b)-1 {
			pattern = `^\+-*\+$`
		} else {
			pattern = `^\|[ *]*\|$`
		}
		if !regexp.MustCompile(pattern).Match(row) {
			return errors.New("invalid board")
		}
		for j, c := range row {
			if c == ' ' {
				for k := 0; k < 9; k++ {
					if b[i-1+k/3][j-1+k%3] == '*' {
						if b[i][j] == ' ' {
							b[i][j] = '1'
							continue
						}
						b[i][j]++
					}
				}
			}
		}
	}
	return nil
}