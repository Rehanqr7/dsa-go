package linkedlist

import . "github.com/Rehanqr7/dsa-go/utils"

/*
Merge Two Sorted Lists (Easy)

Merge two sorted linked lists and return it as a sorted list.

Time Complexity: O(m+n)
Space Complexity: O(1)
*/

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return dummy.Next
}
