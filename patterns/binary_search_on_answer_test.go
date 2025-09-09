package patterns

import "testing"

func TestShipWithinDays(t *testing.T) {
	weights := []int{1,2,3,4,5,6,7,8,9,10}
	if got := ShipWithinDays(weights, 5); got != 15 {
		t.Fatalf("expected 15 got %d", got)
	}
	if got := ShipWithinDays([]int{3,2,2,4,1,4}, 3); got != 6 {
		t.Fatalf("expected 6 got %d", got)
	}
}
