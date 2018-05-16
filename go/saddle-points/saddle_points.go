package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Pair [2]int

type Matrix [][]int

func New(input string) (*Matrix, error) {
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
	return &m, m.validate()
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

func (m Matrix) Saddle() []Pair {
	var saddles []Pair
	for i, row := range m {
		for j, _ := range row {
			if m.isSaddlePoint(i, j) {
				saddles = append(saddles, Pair{i, j})
			}
		}
	}
	return saddles
}

func (m Matrix) isSaddlePoint(i, j int) bool {
	point := m[i][j]
	for _, c := range m[i] {
		if c > point {
			return false
		}
	}
	for _, c := range m.Cols()[j] {
		if c < point {
			return false
		}
	}
	return true
}