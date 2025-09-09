package backtracking

import "testing"

func TestSolveNQueens(t *testing.T) {
	ans := SolveNQueens(4)
	if len(ans) != 2 {
		t.Fatalf("expected 2 solutions got %d", len(ans))
	}
}
