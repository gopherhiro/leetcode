package twopointer

// 5. Longest Palindromic Substring
// 5. 最长回文子串
// 思路：中心拓展法
// time O(N^2) space O(1)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		// 处理回文串长度为奇数的情况：
		// 以 s[i] 为中心的最长回文子串
		si := palindrome(s, i, i)
		// 处理回文串长度为偶数的情况：
		// 以 s[i] 和 s[i+1] 为中心的最长回文子串
		st := palindrome(s, i, i+1)
		// 比较选出最长的回文子串
		if len(si) > len(ans) {
			ans = si
		}
		if len(st) > len(ans) {
			ans = st
		}
	}
	return ans
}

// 在 s 中寻找以 s[left] 和 s[right] 为中心的最长回文串
// 1、当 left == right 时，就是在寻找长度为奇数的回文串。
// 2、当 left != right 时，就是在寻找长度为偶数的回文串。
func palindrome(s string, left, right int) string {
	// 从中心向外扩展
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1 : right]
}
