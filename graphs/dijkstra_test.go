package graphs

import "testing"

func TestDijkstra(t *testing.T) {
	adj := map[int][]Edge{
		0: []Edge{{To:1,W:4},{To:2,W:1}},
		1: []Edge{{To:3,W:1}},
		2: []Edge{{To:1,W:2},{To:3,W:5}},
		3: []Edge{},
	}
	d := Dijkstra(adj, 0)
	if d[3] != 4 { t.Fatalf("shortest 0->3 expected 4 got %d", d[3]) }
}
