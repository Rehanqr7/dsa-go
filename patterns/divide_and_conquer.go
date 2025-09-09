package patterns

/*
Divide & Conquer Pattern
Example: Count Inversions using Merge Sort

Time Complexity: O(n log n)
Space Complexity: O(n)
*/

func CountInversions(nums []int) int {
	_, inv := sortCount(nums)
	return inv
}

func sortCount(a []int) ([]int, int) {
	if len(a) <= 1 { return append([]int(nil), a...), 0 }
	mid := len(a)/2
	left, invL := sortCount(a[:mid])
	right, invR := sortCount(a[mid:])
	merged, cross := mergeCount(left, right)
	return merged, invL + invR + cross
}

func mergeCount(left, right []int) ([]int, int) {
	res := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	inv := 0
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			inv += len(left) - i
			j++
		}
	}
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res, inv
}
