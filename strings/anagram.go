package strings

/*
Valid Anagram (Easy)

Given two strings s and t, return true if t is an anagram of s, and false otherwise.

Time Complexity: O(n)
Space Complexity: O(1) if alphabet is fixed; O(k) otherwise
*/

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int, len(s))
	for _, r := range s {
		count[r]++
	}
	for _, r := range t {
		count[r]--
		if count[r] < 0 {
			return false
		}
	}
	return true
}
