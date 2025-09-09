package stacks

/*
Min Stack (Easy)

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Time Complexity: O(1) per operation
Space Complexity: O(n)
*/

type MinStack struct {
	stack []int
	mins  []int
}

func NewMinStack() *MinStack {
	return &MinStack{stack: make([]int, 0), mins: make([]int, 0)}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	if len(s.mins) == 0 || x <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, x)
	}
}

func (s *MinStack) Pop() {
	if len(s.stack) == 0 {
		return
	}
	x := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	if len(s.mins) > 0 && x == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	if len(s.mins) == 0 {
		return 0
	}
	return s.mins[len(s.mins)-1]
}
