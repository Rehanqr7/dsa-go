package queues

import "testing"

func TestMyQueue(t *testing.T) {
	q := NewMyQueue()
	if !q.Empty() { t.Fatalf("expected empty queue") }
	q.Push(1)
	q.Push(2)
	if q.Peek() != 1 { t.Fatalf("peek expected 1") }
	if q.Pop() != 1 { t.Fatalf("pop expected 1") }
	if q.Empty() { t.Fatalf("expected non-empty queue") }
	if q.Pop() != 2 { t.Fatalf("pop expected 2") }
	if !q.Empty() { t.Fatalf("expected empty queue") }
}
