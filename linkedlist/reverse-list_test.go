package linkedlist

import (
	"reflect"
	"testing"

	"github.com/Rehanqr7/dsa-go/utils"
)

func TestReverseList(t *testing.T) {
	head := utils.BuildLinkedList([]int{1, 2, 3, 4, 5})
	rev := ReverseList(head)
	got := utils.LinkedListToSlice(rev)
	exp := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(got, exp) {
		t.Fatalf("ReverseList failed: got %v want %v", got, exp)
	}
}
