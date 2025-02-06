package main

type Edge struct {
	To     string
	From   string
	Weight int
}

type Graph struct {
	Nodes map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string][]Edge),
	}
}

func (g *Graph) AddEdge(from, to string, weight int) {
	g.Nodes[from] = append(g.Nodes[from], Edge{From: from, To: to, Weight: weight})
	// for an undirected graph
	g.Nodes[to] = append(g.Nodes[to], Edge{From: to, To: from, Weight: weight})
}
