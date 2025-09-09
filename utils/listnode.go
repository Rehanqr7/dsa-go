package utils

// ListNode represents a node in a singly linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildLinkedList builds a linked list from a slice of ints and returns the head.
func BuildLinkedList(values []int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for _, v := range values {
		n := &ListNode{Val: v}
		if head == nil {
			head = n
			tail = n
			continue
		}
		tail.Next = n
		tail = n
	}
	return head
}

// LinkedListToSlice converts a linked list to a slice.
func LinkedListToSlice(head *ListNode) []int {
	result := make([]int, 0)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}
