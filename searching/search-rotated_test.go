package searching

import "testing"

func TestSearchRotated(t *testing.T) {
	nums := []int{4,5,6,7,0,1,2}
	if SearchRotated(nums, 0) != 4 { t.Fatalf("expected index 4") }
	if SearchRotated(nums, 3) != -1 { t.Fatalf("expected -1") }
}
