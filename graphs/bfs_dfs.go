package graphs

/*
Graph Traversals (BFS and DFS) - Easy/Medium

Represent graph as adjacency list: map[int][]int.
Functions return the visitation order starting from a source.
*/

func BFS(adj map[int][]int, start int) []int {
	order := []int{}
	if adj == nil { return order }
	q := []int{start}
	seen := map[int]bool{start: true}
	for len(q) > 0 {
		u := q[0]
		q = q[1:]
		order = append(order, u)
		for _, v := range adj[u] {
			if !seen[v] {
				seen[v] = true
				q = append(q, v)
			}
		}
	}
	return order
}

func DFS(adj map[int][]int, start int) []int {
	order := []int{}
	seen := make(map[int]bool)
	var dfs func(int)
	dfs = func(u int) {
		seen[u] = true
		order = append(order, u)
		for _, v := range adj[u] {
			if !seen[v] {
				dfs(v)
			}
		}
	}
	dfs(start)
	return order
}
