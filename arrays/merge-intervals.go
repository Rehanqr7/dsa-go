package arrays

import "sort"

/*
Merge Intervals - Medium

Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example:
Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]

Time Complexity: O(n log n)
Space Complexity: O(1) extra (excluding output), O(n) including output
*/

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	merged := make([][]int, 0, len(intervals))
	curr := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= curr[1] {
			if intervals[i][1] > curr[1] {
				curr[1] = intervals[i][1]
			}
			continue
		}
		merged = append(merged, curr)
		curr = []int{intervals[i][0], intervals[i][1]}
	}
	merged = append(merged, curr)
	return merged
}
