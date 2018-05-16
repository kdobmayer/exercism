package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden map[string][]string

var plants = map[byte]string{'G': "grass", 'C': "clover", 'R': "radishes", 'V': "violets"}

func NewGarden(diagram string, children []string) (*Garden, error) {
	rows := strings.Split(diagram, "\n")[1:]
	if !isValid(diagram, children, rows) {
		return nil, errors.New("invalid input")
	}

	newc := make([]string, len(children))
	copy(newc, children)
	sort.Strings(newc)

	g := make(Garden)
	for i, c := range newc {
		g[c] = []string{
			plants[rows[0][i*2]],
			plants[rows[0][i*2+1]],
			plants[rows[1][i*2]],
			plants[rows[1][i*2+1]],
		}
	}
	return &g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	p, ok := (*g)[child]
	return p, ok
}

func isValid(diagram string, children, rows []string) bool {
	// check diagram format
	if len(rows) != 2 ||
		len(rows[0]) != len(rows[1]) ||
		len(rows[0])%2 != 0 || len(rows[1])%2 != 0 {
		return false
	}
	// check invalid cup codes
	for _, p := range diagram {
		if p == '\n' {
			continue
		}
		_, ok := plants[byte(p)]
		if !ok {
			return false
		}
	}
	// check duplicate names
	seen := make(map[string]bool)
	for _, c := range children {
		if seen[c] == true {
			return false
		}
		seen[c] = true
	}

	return true
}