package arrays

/*
Sliding Window Maximum - Hard

Given an array nums and an integer k, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Return the max sliding window.

Time Complexity: O(n)
Space Complexity: O(k)
*/

func MaxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 || k == 0 {
		return nil
	}
	// Deque holds indices, maintaining decreasing values
	deque := make([]int, 0, n)
	push := func(i int) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
	}
	pop := func(i int) {
		if len(deque) > 0 && deque[0] <= i-k {
			deque = deque[1:]
		}
	}
	res := make([]int, 0, n-k+1)
	for i := 0; i < n; i++ {
		push(i)
		if i >= k-1 {
			pop(i)
			res = append(res, nums[deque[0]])
		}
	}
	return res
}
