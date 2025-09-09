package heaps

import (
	"reflect"
	"testing"
)

func TestHeapSort(t *testing.T) {
	in := []int{5,1,1,2,0,0}
	exp := []int{0,0,1,1,2,5}
	if got := HeapSort(in); !reflect.DeepEqual(got, exp) {
		t.Fatalf("HeapSort(%v)=%v want %v", in, got, exp)
	}
}
