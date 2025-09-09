package strings

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	cases := []struct {
		s   string
		exp int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}
	for _, c := range cases {
		if got := LengthOfLongestSubstring(c.s); got != c.exp {
			t.Fatalf("LengthOfLongestSubstring(%q)=%d want %d", c.s, got, c.exp)
		}
	}
}
