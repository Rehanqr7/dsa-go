package arrays

/*
Two Sum (Easy)

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

Example:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]

Time Complexity: O(n)
Space Complexity: O(n)
*/

func TwoSum(nums []int, target int) [2]int {
	indexByValue := make(map[int]int)
	for i, v := range nums {
		if j, ok := indexByValue[target-v]; ok {
			return [2]int{j, i}
		}
		indexByValue[v] = i
	}
	return [2]int{-1, -1}
}
