package stacks

/*
Valid Parentheses (Easy)

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

Time Complexity: O(n)
Space Complexity: O(n)
*/

func IsValidParentheses(s string) bool {
	stack := make([]rune, 0, len(s))
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, r := range s {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != pairs[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
