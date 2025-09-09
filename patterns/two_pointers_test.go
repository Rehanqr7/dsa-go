package patterns

import "testing"

func TestHasPairWithSumSorted(t *testing.T) {
	if !HasPairWithSumSorted([]int{1,2,3,4,6}, 6) { t.Fatalf("expected true") }
	if HasPairWithSumSorted([]int{2,5,9,11}, 11) { t.Fatalf("expected false") }
}
