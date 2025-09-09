package dp

import "testing"

func TestLengthOfLIS(t *testing.T) {
	if got := LengthOfLIS([]int{10,9,2,5,3,7,101,18}); got != 4 {
		t.Fatalf("expected 4 got %d", got)
	}
	if got := LengthOfLIS([]int{0,1,0,3,2,3}); got != 4 {
		t.Fatalf("expected 4 got %d", got)
	}
	if got := LengthOfLIS([]int{}); got != 0 {
		t.Fatalf("expected 0 got %d", got)
	}
}
