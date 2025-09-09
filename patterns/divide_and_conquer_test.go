package patterns

import "testing"

func TestCountInversions(t *testing.T) {
	if got := CountInversions([]int{2,4,1,3,5}); got != 3 {
		t.Fatalf("expected 3 got %d", got)
	}
	if got := CountInversions([]int{5,4,3,2,1}); got != 10 {
		t.Fatalf("expected 10 got %d", got)
	}
}
