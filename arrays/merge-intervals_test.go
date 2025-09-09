package arrays

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	cases := []struct {
		in  [][]int
		exp [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}
	for _, c := range cases {
		got := MergeIntervals(c.in)
		if !reflect.DeepEqual(got, c.exp) {
			t.Fatalf("MergeIntervals(%v)=%v want %v", c.in, got, c.exp)
		}
	}
}
