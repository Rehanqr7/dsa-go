package patterns

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	in := [][]int{{1,2},{2,3},{3,4},{1,3}}
	if got := EraseOverlapIntervals(in); got != 1 {
		t.Fatalf("expected 1 got %d", got)
	}
}
