package main

import "fmt"

func main() {
	s := "cbbd"
	r := minInsertions(s)
	fmt.Println(r)
}

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

	memo := generateMemo(n, n, -1)

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

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
	memo := generateMemo(n, n, -1)

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

func generateMemo(m, n, initVal int) [][]int {
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
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

func generateDp(m, n int) [][]int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	return dp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
