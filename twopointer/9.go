package twopointer

import "strconv"

// 9. Palindrome Number
// 9. 回文数
// 思路：int to string, and two pointer
// time O(n) space O(1)
func isPalindromeINT(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	return checkPalindrome(s, 0, len(s)-1)
}
