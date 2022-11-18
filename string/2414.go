package string

// 2414. Length of the Longest Alphabetical Continuous Substring
// 2414. 最长的字母序连续子字符串的长度
// 思路：模拟统计
// time O(n) space O(1)
func longestContinuousSubstring(s string) int {
	ans, count := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			count++
		} else {
			ans = max(ans, count)
			count = 1
		}
	}
	ans = max(ans, count)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
