package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	in := []int{3,6,8,10,1,2,1}
	exp := []int{1,1,2,3,6,8,10}
	QuickSort(in)
	if !reflect.DeepEqual(in, exp) {
		t.Fatalf("QuickSort got %v want %v", in, exp)
	}
}
