package pov

import "fmt"

type Graph map[string]string

func New() *Graph {
	return &Graph{}
}

func (g *Graph) AddNode(nodeLabel string) {
	(*g)[nodeLabel] = ""
}

func (g *Graph) AddArc(from, to string) {
	(*g)[to] = from
}

func (g *Graph) ArcList() []string {
	var arcs []string
	for node, parent := range *g {
		if parent != "" {
			arcs = append(arcs, fmt.Sprintf("%s -> %s", parent, node))
		}
	}
	return arcs
}

func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	changed := New()
	for k, v := range *g {
		(*changed)[k] = v
	}
	for node := newRoot; node != oldRoot; node = (*g)[node] {
		(*changed)[(*g)[node]] = node
	}
	delete(*changed, newRoot)
	return changed
}