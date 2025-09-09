package patterns

import "sort"

/*
Greedy Pattern
Example: Non-overlapping Intervals (Erase Minimum to avoid overlaps)

Given a set of intervals, find the minimum number to remove to make the rest non-overlapping.

Time Complexity: O(n log n)
Space Complexity: O(1) extra
*/

func EraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 { return 0 }
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][1] < intervals[j][1] })
	count := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++ // remove this one
		} else {
			end = intervals[i][1]
		}
	}
	return count
}
