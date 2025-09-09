package searching

import "testing"

func TestBinarySearchAndBounds(t *testing.T) {
	nums := []int{1,2,2,2,3,5}
	if BinarySearch(nums, 3) != 4 { t.Fatalf("binary search failed") }
	if LowerBound(nums, 2) != 1 { t.Fatalf("lower bound failed") }
	if UpperBound(nums, 2) != 4 { t.Fatalf("upper bound failed") }
}
