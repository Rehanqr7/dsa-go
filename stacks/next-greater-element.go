package stacks

/*
Next Greater Element (Monotonic Stack) - Medium

Given an array nums, return an array next where next[i] is the next greater element to the right of nums[i], or -1 if none.

Time Complexity: O(n)
Space Complexity: O(n)
*/

func NextGreaterElements(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := range res { res[i] = -1 }
	stack := make([]int, 0, n) // will store indices with decreasing values
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = nums[i]
		}
		stack = append(stack, i)
	}
	return res
}
