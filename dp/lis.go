package dp

/*
Longest Increasing Subsequence (LIS) Length - Medium

Time Complexity: O(n log n)
Space Complexity: O(n)
*/

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 { return 0 }
	tails := make([]int, 0, len(nums))
	for _, x := range nums {
		// find lower bound of x in tails
		l, r := 0, len(tails)
		for l < r {
			m := (l + r) / 2
			if tails[m] < x { l = m + 1 } else { r = m }
		}
		if l == len(tails) {
			tails = append(tails, x)
		} else {
			tails[l] = x
		}
	}
	return len(tails)
}
