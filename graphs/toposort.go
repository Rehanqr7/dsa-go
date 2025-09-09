package graphs

/*
Topological Sort (Kahn's Algorithm) - Medium

Given a directed acyclic graph (DAG) as adjacency list, return a topological ordering.
If a cycle exists, return an empty slice.
*/

func TopologicalSort(adj map[int][]int) []int {
	in := make(map[int]int)
	for u := range adj { in[u] = 0 }
	for _, nbrs := range adj {
		for _, v := range nbrs { in[v]++ }
	}
	q := make([]int, 0)
	for u := range in {
		if in[u] == 0 { q = append(q, u) }
	}
	order := make([]int, 0, len(in))
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range adj[u] {
			in[v]--
			if in[v] == 0 { q = append(q, v) }
		}
	}
	if len(order) != len(in) { return []int{} }
	return order
}
