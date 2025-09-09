package linkedlist

import . "github.com/Rehanqr7/dsa-go/utils"

/*
Linked List Cycle (Easy)

Given head, determine if the linked list has a cycle in it using O(1) space.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
