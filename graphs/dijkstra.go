package graphs

import (
	"container/heap"
	"math"
)

/*
Dijkstra's Algorithm - Single Source Shortest Paths (non-negative weights)

adj: map[int][]Edge, where Edge{To, W}
Returns distance map from source.
*/

type Edge struct { To int; W int }

func Dijkstra(adj map[int][]Edge, src int) map[int]int {
	dist := make(map[int]int)
	for u := range adj { dist[u] = math.MaxInt/4 }
	dist[src] = 0
	pq := &minPQ{}
	heap.Init(pq)
	heap.Push(pq, item{node: src, dist: 0})
	seen := make(map[int]bool)
	for pq.Len() > 0 {
		it := heap.Pop(pq).(item)
		u := it.node
		if seen[u] { continue }
		seen[u] = true
		for _, e := range adj[u] {
			if nd := dist[u] + e.W; nd < dist[e.To] {
				dist[e.To] = nd
				heap.Push(pq, item{node: e.To, dist: nd})
			}
		}
	}
	return dist
}

type item struct { node int; dist int }

type minPQ []item

func (h minPQ) Len() int { return len(h) }
func (h minPQ) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h minPQ) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minPQ) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *minPQ) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}
