package strings

/*
Palindrome Check (Easy)

Given a string s, return true if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Time Complexity: O(n)
Space Complexity: O(1)
*/

func IsPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isAlnum(s[i]) {
			i++
		}
		for i < j && !isAlnum(s[j]) {
			j--
		}
		if toLower(s[i]) != toLower(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func isAlnum(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}
