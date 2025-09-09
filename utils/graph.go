package utils

type Graph struct {
	Adj map[int][]int
}

func NewGraph() *Graph {
	return &Graph{Adj: make(map[int][]int)}
}

func (g *Graph) AddEdge(u, v int, directed bool) {
	g.Adj[u] = append(g.Adj[u], v)
	if !directed {
		g.Adj[v] = append(g.Adj[v], u)
	}
}

func (g *Graph) Nodes() []int {
	nodes := make([]int, 0, len(g.Adj))
	for u := range g.Adj {
		nodes = append(nodes, u)
	}
	return nodes
}
