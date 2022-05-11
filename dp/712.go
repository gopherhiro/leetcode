package dp

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
