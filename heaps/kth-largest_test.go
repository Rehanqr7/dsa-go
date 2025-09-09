package heaps

import "testing"

func TestFindKthLargest(t *testing.T) {
	if got := FindKthLargest([]int{3,2,1,5,6,4}, 2); got != 5 {
		t.Fatalf("expected 5 got %d", got)
	}
	if got := FindKthLargest([]int{3,2,3,1,2,4,5,5,6}, 4); got != 4 {
		t.Fatalf("expected 4 got %d", got)
	}
}
