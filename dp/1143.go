package dp

// 1143. Longest Common Subsequence
// 1143. 最长公共子序列
// 思路：Dynamic Programming
// time O(m*n) space O(m*n)
// 定义：s1[0..i-1] 和 s2[0..j-1] 的 lcs 长度为 dp[i][j]
// 目标：s1[0..m-1] 和 s2[0..n-1] 的 lcs 长度，即 dp[m][n]
// 技巧：根据字符串s1,s2 画出其二维 dp 数组
// dp[0][...] 和 dp[...][0] 初始化为0
func longestCommonSubsequence(s1, s2 string) int {
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return 0
	}
	// base case: dp[0][..] = dp[..][0] = 0
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// i,j 从0开始
	for i, c1 := range s1 {
		for j, c2 := range s2 {
			if c1 == c2 {
				dp[i+1][j+1] = 1 + dp[i][j]
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	return dp[m][n]
}

// 1143. Longest Common Subsequence
// 1143. 最长公共子序列
// 思路：Recursion + Memoization
// time O(m*n) space O(m*n)
// 技巧：画出递归调用树，你会发现重叠子问题。
// 从而可以使用备忘录进行优化
func longestCommonSubsequenceO(s1, s2 string) int {
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return 0
	}
	// 备忘录值为 -1 代表未曾计算过
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}

	// helper(i,j) 表示 s1[i..] 和 s2[j..] 的最长公共子序列长度
	var helper func(i, j int) int
	helper = func(i, j int) int {
		// base case
		if i == len(s1) || j == len(s2) {
			return 0
		}
		// 如果之前计算过，则直接返回备忘录中的答案
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		// 根据 s1[i] 和 s2[j] 情况， 做选择
		if s1[i] == s2[j] {
			memo[i][j] = 1 + helper(i+1, j+1)
		} else {
			memo[i][j] = max(helper(i, j+1), helper(i+1, j))
		}
		return memo[i][j]
	}
	return helper(0, 0)
}

// 1143. Longest Common Subsequence
// 1143. 最长公共子序列
// 思路：Recursion
// time O(2^n) space O(2^n)
// 技巧：画出递归调用树
func longestCommonSubsequenceR(s1, s2 string) int {
	m, n := len(s1), len(s2)
	if m == 0 || n == 0 {
		return 0
	}
	// helper(i,j) 表示 s1[i..] 和 s2[j..] 的最长公共子序列长度
	var helper func(i, j int) int
	helper = func(i, j int) int {
		// base case
		if i == m || j == n {
			return 0
		}
		if s1[i] == s2[j] {
			return 1 + helper(i+1, j+1)
		}
		return max(helper(i, j+1), helper(i+1, j))
	}
	return helper(0, 0)
}
