package arrays

import "testing"

func TestMaximumSubarray(t *testing.T) {
	cases := []struct {
		in  []int
		exp int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{-1, -2, -3}, -1},
	}
	for _, c := range cases {
		if got := MaximumSubarray(c.in); got != c.exp {
			t.Fatalf("MaximumSubarray(%v) = %d, want %d", c.in, got, c.exp)
		}
	}
}
