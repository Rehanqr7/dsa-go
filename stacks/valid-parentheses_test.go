package stacks

import "testing"

func TestIsValidParentheses(t *testing.T) {
	cases := []struct{
		in string
		exp bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}
	for _, c := range cases {
		if got := IsValidParentheses(c.in); got != c.exp {
			t.Fatalf("IsValidParentheses(%q)=%v want %v", c.in, got, c.exp)
		}
	}
}
