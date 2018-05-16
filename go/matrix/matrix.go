package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(input string) (Matrix, error) {
	rows := strings.Split(input, "\n")
	m := make(Matrix, len(rows))
	for i, r := range rows {
		cols := strings.Split(strings.TrimSpace(r), " ")
		m[i] = make([]int, len(cols))
		for j, c := range cols {
			v, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			m[i][j] = v
		}
	}
	return m, m.validate()
}

func (m Matrix) validate() error {
	length := 0
	for _, row := range m {
		if length == 0 {
			length = len(row)
		}
		if len(row) != length {
			return errors.New("Invalid matrix")
		}
	}
	return nil
}

func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]int, len(m[i]))
		for j := 0; j < len(rows[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}
	return rows
}

func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := 0; i < len(cols); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(cols[i]); j++ {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

func (m Matrix) Set(r, c, val int) (ok bool) {
	if ok = r >= 0 && c >= 0 && r < len(m) && c < len(m[0]); ok {
		m[r][c] = val
	}
	return
}