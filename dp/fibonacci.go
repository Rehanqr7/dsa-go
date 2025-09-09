package dp

/*
Fibonacci (Easy)
Compute the nth Fibonacci number.

Time Complexity: O(n)
Space Complexity: O(1) iterative
*/

func Fib(n int) int {
	if n <= 1 { return n }
	prev2, prev1 := 0, 1
	for i := 2; i <= n; i++ {
		prev2, prev1 = prev1, prev1+prev2
	}
	return prev1
}
