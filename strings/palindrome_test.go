package strings

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		in  string
		exp bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{"", true},
	}
	for _, c := range cases {
		if got := IsPalindrome(c.in); got != c.exp {
			t.Fatalf("IsPalindrome(%q)=%v want %v", c.in, got, c.exp)
		}
	}
}
