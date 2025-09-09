package strings

/*
Longest Substring Without Repeating Characters (Medium)

Given a string s, find the length of the longest substring without repeating characters.

Time Complexity: O(n)
Space Complexity: O(min(n, alphabet))
*/

func LengthOfLongestSubstring(s string) int {
	last := make(map[rune]int)
	start := 0
	best := 0
	for i, r := range []rune(s) {
		if j, ok := last[r]; ok && j >= start {
			start = j + 1
		}
		if i-start+1 > best {
			best = i - start + 1
		}
		last[r] = i
	}
	return best
}
