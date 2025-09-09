package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	in := []int{5, 2, 3, 1}
	exp := []int{1, 2, 3, 5}
	if got := MergeSort(in); !reflect.DeepEqual(got, exp) {
		t.Fatalf("MergeSort got %v want %v", got, exp)
	}
}
