package memoization

import "strings"

// 91. Decode Ways
// 91. 解码方法
// 思路：dp
// time O(n) space O(n)
func numDecodingsD(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1
	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}
		if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && strings.ContainsRune("0123456", rune(s[i+1])))) {
			dp[i] += dp[i+2]
		}
	}
	return dp[0]
}

// 91. Decode Ways
// 91. 解码方法
// 思路：memory search
// time O(n) space O(n)
func numDecodings(s string) int {
	n := len(s)
	m := make(map[int]int, n+1)
	var dfs func(i int) int
	dfs = func(i int) int {
		// base case
		// 达到字符串末尾时结束
		// 找到了一条路径，所以返回 1
		if i == len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		// query from memory
		if count, ok := m[i]; ok {
			return count
		}
		// return dfs(i+1) + dfs(i+2)
		// dfs(i+1) don't need constraint condition
		// dfs(i+2) need constraint condition, because two digit number
		// must less than or equal to 26
		ans := dfs(i + 1)
		if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && strings.ContainsRune("0123456", rune(s[i+1])))) {
			ans += dfs(i + 2)
		}
		m[i] = ans
		return ans
	}
	return dfs(0)
}

// 91. Decode Ways
// 91. 解码方法
// 思路：recursion
// time O(2^n) space O(n)
// TLE
func numDecodingsR(s string) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		// base case
		if i == len(s) {
			return 1
		}
		if s[i] == '0' {
			return 0
		}
		// return dfs(i+1) + dfs(i+2)
		// dfs(i+1) don't need constraint condition
		// dfs(i+2) need constraint condition, because two digit number
		// must less than or equal to 26
		ans := dfs(i + 1)
		if i+1 < len(s) && (s[i] == '1' || (s[i] == '2' && strings.ContainsRune("0123456", rune(s[i+1])))) {
			ans += dfs(i + 2)
		}
		return ans
	}
	return dfs(0)
}
