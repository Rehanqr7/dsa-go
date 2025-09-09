package sorting

/*
QuickSort (Medium)

Time Complexity: Average O(n log n), Worst O(n^2)
Space Complexity: O(log n) recursion
*/

func QuickSort(nums []int) {
	if len(nums) <= 1 { return }
	qs(nums, 0, len(nums)-1)
}

func qs(a []int, l, r int) {
	if l >= r { return }
	p := partition(a, l, r)
	qs(a, l, p-1)
	qs(a, p+1, r)
}

func partition(a []int, l, r int) int {
	pivot := a[r]
	i := l
	for j := l; j < r; j++ {
		if a[j] < pivot {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	return i
}
