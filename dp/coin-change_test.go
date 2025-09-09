package dp

import "testing"

func TestCoinChange(t *testing.T) {
	if got := CoinChange([]int{1,2,5}, 11); got != 3 {
		t.Fatalf("expected 3 got %d", got)
	}
	if got := CoinChange([]int{2}, 3); got != -1 {
		t.Fatalf("expected -1 got %d", got)
	}
}
