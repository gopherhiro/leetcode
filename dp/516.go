package dp

// 516. Longest Palindromic Subsequence
// 516. 最长回文子序列
// 思路：recursion + memo
// 字符串:s[i...i+1......j-1...j]
// dp函数定义：在子串s[i..j]，最长回文子序列的长度为 dp(i,j)
func longestPalindromeSubseqM(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	// 备忘录
	memo := genmemo(n, n, -1)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		// 查备忘录
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		if s[i] == s[j] {
			memo[i][j] = dp(i+1, j-1) + 2
		} else {
			memo[i][j] = max(dp(i, j-1), dp(i+1, j))
		}
		return memo[i][j]
	}
	return dp(0, n-1)
}

// 516. Longest Palindromic Subsequence
// 516. 最长回文子序列
// 思路：recursion
// 字符串:s[i...i+1......j-1...j]
// dp函数定义：在子串s[i..j]，最长回文子序列的长度为 dp(i,j)
func longestPalindromeSubseqR(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i > j {
			return 0
		}
		if i == j {
			return 1
		}

		if s[i] == s[j] {
			return dp(i+1, j-1) + 2
		}
		return max(dp(i, j-1), dp(i+1, j))
	}
	return dp(0, n-1)
}

// 516. Longest Palindromic Subsequence
// 516. 最长回文子序列
// 思路：dynamic programming
// 字符串:s[i...i+1......j-1...j]
// dp数组定义：在子串s[i..j]，最长回文子序列的长度为 dp[i][j]
// time O(n^2) space O(n^2)
func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	// s[i..j] 与 dp[i][j] 一一对应

	// base case
	// i > j , dp[i][j] = 0, 即空串，则回文串长度为 0
	dp := generateDp(n, n)

	// i == j , dp[i][j] = 1, 即只有一个字符，则回文串长度为 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// 从下至上，从左至右遍历（即斜向上遍历）
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				// 字符串:s[i...i+1......j-1...j]
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}
