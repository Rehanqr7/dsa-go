package strings

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s, t string
		exp  bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}
	for _, c := range cases {
		if got := IsAnagram(c.s, c.t); got != c.exp {
			t.Fatalf("IsAnagram(%q,%q)=%v want %v", c.s, c.t, got, c.exp)
		}
	}
}
