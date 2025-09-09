package heaps

import "container/heap"

/*
Kth Largest Element in an Array (Medium)

Time Complexity: O(n log k)
Space Complexity: O(k)
*/

func FindKthLargest(nums []int, k int) int {
	h := &minHeap{}
	heap.Init(h)
	for _, v := range nums {
		if h.Len() < k {
			heap.Push(h, v)
		} else if v > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}
	return (*h)[0]
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}
