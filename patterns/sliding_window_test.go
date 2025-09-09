package patterns

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	if got := MinSubArrayLen(7, []int{2,3,1,2,4,3}); got != 2 {
		t.Fatalf("expected 2 got %d", got)
	}
	if got := MinSubArrayLen(4, []int{1,4,4}); got != 1 {
		t.Fatalf("expected 1 got %d", got)
	}
	if got := MinSubArrayLen(11, []int{1,1,1,1,1,1,1,1}); got != 0 {
		t.Fatalf("expected 0 got %d", got)
	}
}
