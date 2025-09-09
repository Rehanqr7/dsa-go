package queues

/*
Implement Queue using Stacks (Easy)

Implement a first in first out (FIFO) queue using only two stacks.

Time Complexity: Amortized O(1) per operation
Space Complexity: O(n)
*/

type MyQueue struct {
	in  []int
	out []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{in: make([]int, 0), out: make([]int, 0)}
}

func (q *MyQueue) Push(x int) {
	q.in = append(q.in, x)
}

func (q *MyQueue) Pop() int {
	q.move()
	if len(q.out) == 0 { return 0 }
	x := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return x
}

func (q *MyQueue) Peek() int {
	q.move()
	if len(q.out) == 0 { return 0 }
	return q.out[len(q.out)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.in) == 0 && len(q.out) == 0
}

func (q *MyQueue) move() {
	if len(q.out) == 0 {
		for len(q.in) > 0 {
			q.out = append(q.out, q.in[len(q.in)-1])
			q.in = q.in[:len(q.in)-1]
		}
	}
}
