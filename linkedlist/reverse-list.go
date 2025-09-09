package linkedlist

/*
Reverse Linked List (Easy)

Given the head of a singly linked list, reverse the list, and return the reversed list.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
