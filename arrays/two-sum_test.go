package arrays

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		expect [2]int
	}{
		{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		{[]int{3, 2, 4}, 6, [2]int{1, 2}},
		{[]int{3, 3}, 6, [2]int{0, 1}},
	}
	for _, c := range cases {
		got := TwoSum(c.nums, c.target)
		if (got != c.expect) && !(c.nums[got[0]]+c.nums[got[1]] == c.target) {
			t.Fatalf("TwoSum(%v,%d) = %v, expected indices summing to target", c.nums, c.target, got)
		}
	}
}
