package tree

import (
	"sort"
)

var nodeCount int

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

type Mismatch struct{}

func (m Mismatch) Error() string {
	return "c"
}

func isValid(records []Record) bool {
	for i, r := range records {
		if r.ID != i || r.ID < r.Parent {
			return false
		}
	}
	return true
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })
	if !isValid(records) {
		return nil, Mismatch{}
	}
	nodeCount = 0
	n := NewNode(records, 0)
	if nodeCount != len(records) {
		return nil, Mismatch{}
	}
	return n, nil
}

func NewNode(records []Record, id int) *Node {
	nodeCount++
	n := &Node{ID: id}
	for i := id + 1; i < len(records); i++ {
		if records[i].Parent == id {
			n.Children = append(n.Children, NewNode(records, i))
		}
	}
	return n
}