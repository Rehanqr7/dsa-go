package heaps

import "container/heap"

/*
Heap Sort (Medium)

Time Complexity: O(n log n)
Space Complexity: O(n) using heap structure
*/

func HeapSort(nums []int) []int {
	h := &maxHeap{}
	heap.Init(h)
	for _, v := range nums { heap.Push(h, v) }
	res := make([]int, 0, len(nums))
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(int))
	}
	// res contains elements from largest to smallest; reverse to get ascending
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	v := old[n-1]
	*h = old[:n-1]
	return v
}
