package slidingwindow

// 424. Longest Repeating Character Replacement
// 424. 替换后的最长重复字符
// solution:sliding windows
// time O(n) space O(1)
func characterReplacement(s string, k int) int {
	res := 0
	count := [26]int{} // the number of character
	left, right, maxCount := 0, 0, 0
	for right < len(s) {
		// largest count of a single, unique
		// character in the current window
		count[s[right]-'A']++
		maxCount = max(maxCount, count[s[right]-'A'])

		// there are more characters in the window than we can replace (k)
		// need to shrink the window.
		for right-left+1-maxCount > k {
			count[s[left]-'A']--
			left++
		}

		res = max(res, right-left+1)
		right++
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
