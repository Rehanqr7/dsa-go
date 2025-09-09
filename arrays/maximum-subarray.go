package arrays

/*
Maximum Subarray (Kadane's Algorithm) - Medium

Given an integer array nums, find the subarray with the largest sum, and return its sum.

Example:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func MaximumSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	best := nums[0]
	curr := nums[0]
	for i := 1; i < len(nums); i++ {
		if curr < 0 {
			curr = nums[i]
		} else {
			curr += nums[i]
		}
		if curr > best {
			best = curr
		}
	}
	return best
}
