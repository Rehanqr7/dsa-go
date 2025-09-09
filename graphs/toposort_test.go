package graphs

import (
	"testing"
)

func TestTopologicalSort(t *testing.T) {
	adj := map[int][]int{
		5: {2,0},
		4: {0,1},
		2: {3},
		3: {1},
		0: {},
		1: {},
	}
	order := TopologicalSort(adj)
	if len(order) != 6 { t.Fatalf("expected 6 nodes got %d", len(order)) }
	pos := map[int]int{}
	for i, v := range order { pos[v] = i }
	// verify edges go from earlier to later
	for u, nbrs := range adj {
		for _, v := range nbrs {
			if !(pos[u] < pos[v]) {
				t.Fatalf("edge %d->%d violates topological order", u, v)
			}
		}
	}
}
