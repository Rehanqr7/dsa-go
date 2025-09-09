package backtracking

/*
N-Queens (Hard)
Return all distinct solutions to the n-queens puzzle.

Time Complexity: O(n!)
Space Complexity: O(n)
*/

func SolveNQueens(n int) [][]string {
	cols := make(map[int]bool)
	d1 := make(map[int]bool) // r - c
	d2 := make(map[int]bool) // r + c
	board := make([]int, n) // board[r] = c
	res := [][]string{}
	var place func(int)
	place = func(r int) {
		if r == n {
			res = append(res, render(board))
			return
		}
		for c := 0; c < n; c++ {
			if cols[c] || d1[r-c] || d2[r+c] { continue }
			cols[c], d1[r-c], d2[r+c] = true, true, true
			board[r] = c
			place(r+1)
			cols[c], d1[r-c], d2[r+c] = false, false, false
		}
	}
	place(0)
	return res
}

func render(board []int) []string {
	n := len(board)
	ans := make([]string, n)
	for r := 0; r < n; r++ {
		row := make([]byte, n)
		for c := 0; c < n; c++ { row[c] = '.' }
		row[board[r]] = 'Q'
		ans[r] = string(row)
	}
	return ans
}
