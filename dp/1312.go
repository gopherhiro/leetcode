package dp

// 1312. Minimum Insertion Steps to Make a String Palindrome
// 1312. 让字符串成为回文串的最少插入次数
// 思路：dynamic programming
// dp 数组定义：子串 s[i..j]变为回文串的「最少操作次数」为 dp[i][j]
// time O(n^2) space O(n^2)
func minInsertions(s string) int {
	n := len(s)
	if n <= 1 {
		return 0 // 空串 or 一个字符，本身具有回文性，不需要插入
	}
	dp := generateDp(n, n)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// s[i] == s[j], no need operation, just skip it
				dp[i][j] = dp[i+1][j-1]
			} else {
				// s[i] != s[j], choose min of dp(i, j-1) and dp(i+1, j)
				// plus one
				dp[i][j] = min(dp[i][j-1], dp[i+1][j]) + 1
			}
		}
	}
	return dp[0][n-1]
}

// 1312. Minimum Insertion Steps to Make a String Palindrome
// 1312. 让字符串成为回文串的最少插入次数
// 思路：记忆化递归
// dp 函数定义：子串 s[i..j]变为回文串的「最少操作次数」为 dp(i,j)
// time O(n^2) space O(n^2)
func minInsertionsM(s string) int {
	n := len(s)
	if n <= 1 {
		return 0 // 空串 or 一个字符，本身具有回文性，不需要插入
	}

	memo := genmemo(n, n, -1)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i > j {
			return 0
		}
		if i == j {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s[i] == s[j] {
			// s[i] == s[j], no operation, just skip it
			memo[i][j] = dp(i+1, j-1)
		} else {
			// s[i] != s[j], choose min of dp(i, j-1) and dp(i+1, j)
			// plus one
			memo[i][j] = min(dp(i, j-1), dp(i+1, j)) + 1
		}
		return memo[i][j]
	}
	return dp(0, n-1)
}

// 1312. Minimum Insertion Steps to Make a String Palindrome
// 1312. 让字符串成为回文串的最少插入次数
// 思路：recursion
// dp 函数定义：子串 s[i..j]变为回文串的「最少操作次数」为 dp(i,j)
// time O(2^n) space O(2^n)
func minInsertionsR(s string) int {
	n := len(s)
	if n <= 1 {
		return 0 // 空串 or 一个字符，本身具有回文性，不需要插入
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i > j {
			return 0
		}
		if i == j {
			return 0
		}
		if s[i] == s[j] {
			// s[i] == s[j], no operation, just skip it
			return dp(i+1, j-1)
		}
		// s[i] != s[j], choose min of dp(i, j-1) and dp(i+1, j)
		// plus one
		return min(dp(i, j-1), dp(i+1, j)) + 1
	}
	return dp(0, n-1)
}
