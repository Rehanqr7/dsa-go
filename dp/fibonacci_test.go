package dp

import "testing"

func TestFib(t *testing.T) {
	cases := []struct{ n, exp int }{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {10, 55},
	}
	for _, c := range cases {
		if got := Fib(c.n); got != c.exp {
			t.Fatalf("Fib(%d)=%d want %d", c.n, got, c.exp)
		}
	}
}
