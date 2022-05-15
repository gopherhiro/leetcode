package twopointer

// 680. Valid Palindrome II
// 680. 验证回文字符串 Ⅱ
// 思路：
// 1、先验证 s 是否是回文串
// 2、如果 s 是回文串，则return true
// 3、如果 s 不是回文串，则把 s 分成 2 份。
// 即：[left..right-1] 和 [left+1..right] 分别验证回文性。
// time O(N), space O(1)
func validPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			// 去掉一个字符情况分为：[left..right-1] 和 [left+1..right] 区域
			// 分别 check 以上两种情况即可。
			return checkPalindrome(s, left, right-1) || checkPalindrome(s, left+1, right)
		}
		left++
		right--
	}
	return true
}
