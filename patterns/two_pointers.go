package patterns

/*
Two Pointers Pattern
Example: Pair Sum in Sorted Array

Given a sorted array nums and a target, determine if there exists a pair that sums to target.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func HasPairWithSumSorted(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l < r {
		s := nums[l] + nums[r]
		if s == target { return true }
		if s < target { l++ } else { r-- }
	}
	return false
}
