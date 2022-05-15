package twopointer

import (
	"regexp"
	"strings"
)

// 125. Valid Palindrome
// 125. 验证回文串
// 思路：two pointer
// 1、去掉所有非数字、字母字符。
// 2、统一转为 lower case or upper case
// 3、判断回文
// time O(N) space O(1)
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 去掉所有非数字、字母字符。
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = re.ReplaceAllString(s, "")
	// to lower case
	s = strings.ToLower(s)
	// two pointer 判断是否回文
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 125. Valid Palindrome
// 125. 验证回文串
// 思路：在原字符串上使用 two pointer
// time O(N) space O(1)
func isPalindromeO(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)

	left, right := 0, len(s)-1
	for left < right {
		// find left alphanumeric characters
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}

		// find right alphanumeric characters
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}

		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// 判断是否是字符或数字
func isAlphanumeric(ch byte) bool {
	return ch >= 'a' && ch <= 'z' ||
		ch >= 'A' && ch <= 'Z' ||
		ch >= '0' && ch <= '9'
}
