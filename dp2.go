package main

import "fmt"

func main() {
	s1 := "horse"
	s2 := "ros"
	r := minDistance(s1, s2)
	fmt.Println(r)
}

// 72. Edit Distance
// 72. 编辑距离
// 思路：DP(Dynamic Programming)
// time O(mn) space O(mn)
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := genDP(m, n)

	// base case
	// dp[i][0] = i
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	// dp[0][j] = j
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// word1[i] != word2[j]
				// have three options:insert,delete,replace,
				// choose min
				insert := dp[i][j-1] + 1    // insert 之后的操作
				delete := dp[i-1][j] + 1    // delete 之后的操作
				replace := dp[i-1][j-1] + 1 // replace 之后的操作
				dp[i][j] = min(insert, min(delete, replace))
			}
		}
	}
	return dp[m][n]
}

func genDP(m, n int) [][]int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	return dp
}

// 72. Edit Distance
// 72. 编辑距离（遍历 word1,word2 从前往后）
// 思路：Recursion + Memo
// time O(mn) space O(mn)
func minDistanceO(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	memo := genMemo(m, n, -1)

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == m {
			return n - j
		}
		if j == n {
			return m - i
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			memo[i][j] = dp(i+1, j+1)
		} else {
			// if word1[i] != word2[j]
			// have three options： insert,delete,replace,
			// choose min
			insert := dp(i, j+1) + 1    // insert 之后的操作
			delete := dp(i+1, j) + 1    // delete 之后的操作
			replace := dp(i+1, j+1) + 1 // replace 之后的操作
			memo[i][j] = min(insert, min(delete, replace))
		}
		return memo[i][j]
	}
	return dp(0, 0)
}

// 72. Edit Distance
// 72. 编辑距离（遍历 word1,word2 从前往后）
// 思路：Recursion
// time O(2^n) space O(2^n)
func minDistanceR(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == m {
			return n - j
		}
		if j == n {
			return m - i
		}
		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			return dp(i+1, j+1)
		}
		// if word1[i] != word2[j]
		// have three options： insert,delete,replace,
		// choose min
		insert := dp(i, j+1) + 1    // insert 之后的操作
		delete := dp(i+1, j) + 1    // delete 之后的操作
		replace := dp(i+1, j+1) + 1 // replace 之后的操作
		return min(insert, min(delete, replace))
	}
	return dp(0, 0)
}

// 72. Edit Distance
// 72. 编辑距离
// 思路：Recursion + Memo （遍历 word1,word2 从后往前）
// time O(m*n) space O(m*n)
func minDistanceM(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 备忘录
	memo := genMemo(m, n, -1)

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}

		// 查询备忘录
		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			memo[i][j] = dp(i-1, j-1)
		} else {
			// if word1[i] != word2[j] have three options： insert,delete,replace,
			// choose min
			insert := dp(i, j-1) + 1    // insert 之后的操作
			delete := dp(i-1, j) + 1    // delete 之后的操作
			replace := dp(i-1, j-1) + 1 // replace 之后的操作
			memo[i][j] = min(insert, min(delete, replace))
		}
		return memo[i][j]
	}
	return dp(m-1, n-1)
}

func genMemo(m, n, initVal int) [][]int {
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}

// 72. Edit Distance
// 72. 编辑距离
// 思路：Recursion（遍历 word1,word2 从后往前）
// time O(2^n) space O(2^n)
func minDistanceR1(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	// 遍历 word1,word2 从后往前
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i < 0 {
			return j + 1
		}
		if j < 0 {
			return i + 1
		}
		// if word1[i] == word2[j]
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		}
		// if word1[i] != word2[j] have three options： insert,delete,replace,
		// choose min
		insert := dp(i, j-1) + 1    // insert 之后的操作
		delete := dp(i-1, j) + 1    // delete 之后的操作
		replace := dp(i-1, j-1) + 1 // replace 之后的操作
		return min(insert, min(delete, replace))
	}

	return dp(m-1, n-1)
}

// 712. Minimum ASCII Delete Sum for Two Strings
// 712. 两个字符串的最小ASCII删除和
// 思路：Dynamic Programming
// time O(m*n) space O(m*n)
func minimumDeleteSumD(a string, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	// base case
	// init dp[...][0]
	for i := 0; i < m; i++ {
		dp[i+1][0] = dp[i][0] + int(a[i])
	}
	// init dp[0][...]
	for j := 0; j < n; j++ {
		dp[0][j+1] = dp[0][j] + int(b[j])
	}

	for i, ca := range a {
		for j, cb := range b {
			if ca == cb {
				dp[i+1][j+1] = dp[i][j]
			} else {
				opt1 := dp[i][j+1] + int(a[i]) // 移动 a，故 +a[i](删除该字符)
				opt2 := dp[i+1][j] + int(b[j]) // 移动 b，故 +b[j](删除该字符)
				dp[i+1][j+1] = min(opt1, opt2)
			}
		}
	}
	return dp[m][n]
}

// 712. Minimum ASCII Delete Sum for Two Strings
// 712. 两个字符串的最小ASCII删除和
// 思路：Recursion + Memoization
// time O(m*n) space O(m*n)
func minimumDeleteSumM(a string, b string) int {
	m, n := len(a), len(b)
	// 求从 start 到 end 的所有字符ascii的和
	var sumToEnd func(s string, start int) int
	sumToEnd = func(s string, start int) int {
		sum := 0
		for i := start; i < len(s); i++ {
			sum += int(s[i])
		}
		return sum
	}

	// 备忘录值为 -1 代表未曾计算过
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = -1
		}
	}

	var helper func(i, j int) int
	helper = func(i, j int) int {
		// base case
		if i == m {
			return sumToEnd(b, j)
		}
		if j == n {
			return sumToEnd(a, i)
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// 相等，表示是公共子序列，不用删除，跳过。
		if a[i] == b[j] {
			memo[i][j] = helper(i+1, j+1)
		} else {
			// 不相等的话，则枚举 2 option 中的min
			opt1 := helper(i+1, j) + int(a[i])
			opt2 := helper(i, j+1) + int(b[j])
			memo[i][j] = min(opt1, opt2)
		}
		return memo[i][j]
	}
	return helper(0, 0)
}

// 712. Minimum ASCII Delete Sum for Two Strings
// 712. 两个字符串的最小ASCII删除和
// 思路：Force Recursion
// time O(2^n) space O(2^n)
// 备注：超出时间限制
func minimumDeleteSumR(a string, b string) int {
	m, n := len(a), len(b)
	// 求从 start 到 end 的所有字符ascii的和
	var sumToEnd func(s string, start int) int
	sumToEnd = func(s string, start int) int {
		sum := 0
		for i := start; i < len(s); i++ {
			sum += int(s[i])
		}
		return sum
	}

	var helper func(i, j int) int
	helper = func(i, j int) int {
		// base case
		if i == m {
			return sumToEnd(b, j)
		}
		if j == n {
			return sumToEnd(a, i)
		}

		// 相等，表示是公共子序列，不用删除，跳过。
		if a[i] == b[j] {
			return helper(i+1, j+1)
		}
		// 不相等的话，则枚举 2 option 中的min
		return min(helper(i+1, j)+int(a[i]), helper(i, j+1)+int(b[j]))
	}
	return helper(0, 0)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 583. Delete Operation for Two Strings
// 583. 两个字符串的删除操作
// 思路：两个字符串的长度分别减去[最长公共子序列]长度，再求和即可。
// time O(m*n) space O(m*n)
func minDeleteStep(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	lcs := longestCommonSubsequence(word1, word2)
	// 步数即是：两个单词变为其最长公共子序列过程，所需要删除的字符个数
	return m - lcs + n - lcs
}

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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
