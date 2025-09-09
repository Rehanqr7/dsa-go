package graphs

import (
	"reflect"
	"testing"
)

func TestBFS_DFS(t *testing.T) {
	adj := map[int][]int{
		0: []int{1,2},
		1: []int{2},
		2: []int{0,3},
		3: []int{3},
	}
	b := BFS(adj, 2)
	// BFS order can vary if adjacency differs; here it's deterministic by insertion
	if !reflect.DeepEqual(b, []int{2,0,3,1}) {
		t.Fatalf("BFS got %v", b)
	}
	d := DFS(adj, 2)
	if !reflect.DeepEqual(d, []int{2,0,1,3}) {
		t.Fatalf("DFS got %v", d)
	}
}
