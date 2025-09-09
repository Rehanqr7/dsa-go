package arrays

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		exp  []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
	}
	for _, c := range cases {
		if got := MaxSlidingWindow(c.nums, c.k); !reflect.DeepEqual(got, c.exp) {
			t.Fatalf("MaxSlidingWindow(%v,%d)=%v want %v", c.nums, c.k, got, c.exp)
		}
	}
}
