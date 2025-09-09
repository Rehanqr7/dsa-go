package sorting

/*
MergeSort (Medium)

Time Complexity: O(n log n)
Space Complexity: O(n)
*/

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return append([]int(nil), nums...)
	}
	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	return merge(left, right)
}

func merge(a, b []int) []int {
	res := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			j++
		}
	}
	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res
}
