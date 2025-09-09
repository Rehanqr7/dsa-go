package searching

/*
Binary Search Variants (Easy/Medium)
*/

func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target { return m }
		if nums[m] < target { l = m + 1 } else { r = m - 1 }
	}
	return -1
}

func LowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] < target { l = m + 1 } else { r = m }
	}
	return l
}

func UpperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] <= target { l = m + 1 } else { r = m }
	}
	return l
}
