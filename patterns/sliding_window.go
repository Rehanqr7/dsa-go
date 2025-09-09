package patterns

/*
Sliding Window Pattern
Example: Minimum Size Subarray Sum (>= target)

Given an array of positive integers nums and a positive integer target, return the minimal length of a
contiguous subarray of which the sum is greater than or equal to target. If there is no such subarray, return 0.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func MinSubArrayLen(target int, nums []int) int {
	n := len(nums)
	best := n + 1
	sum := 0
	left := 0
	for right := 0; right < n; right++ {
		sum += nums[right]
		for sum >= target {
			if length := right - left + 1; length < best {
				best = length
			}
			sum -= nums[left]
			left++
		}
	}
	if best == n+1 { return 0 }
	return best
}
