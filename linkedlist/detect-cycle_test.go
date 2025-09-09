package linkedlist

import (
	"testing"

	"github.com/Rehanqr7/dsa-go/utils"
)

func TestHasCycle(t *testing.T) {
	head := utils.BuildLinkedList([]int{3, 2, 0, -4})
	// create cycle: tail -> node with value 2 (index 1)
	tail := head
	var join *utils.ListNode
	idx := 0
	for tail.Next != nil {
		if idx == 1 {
			join = tail
		}
		tail = tail.Next
		idx++
	}
	tail.Next = join
	if !HasCycle(head) {
		t.Fatalf("expected cycle to be detected")
	}
	// break cycle and test false
	tail.Next = nil
	if HasCycle(head) {
		t.Fatalf("expected no cycle")
	}
}
