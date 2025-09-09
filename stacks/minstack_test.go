package stacks

import "testing"

func TestMinStack(t *testing.T) {
	s := NewMinStack()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	if s.GetMin() != -3 { t.Fatalf("GetMin expected -3") }
	s.Pop()
	if s.Top() != 0 { t.Fatalf("Top expected 0") }
	if s.GetMin() != -2 { t.Fatalf("GetMin expected -2") }
}
