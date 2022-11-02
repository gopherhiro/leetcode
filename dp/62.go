package dp

// 62. Unique Paths
// 62. 不同路径
// 思路：dfs
// time O(n * 2^n) space O(n)
// Time Limit Exceeded
func uniquePathsR(m int, n int) int {
	// counter: the number of paths of from [i,j] to [m-1, n-1]
	var counter func(i, j int) int
	counter = func(i, j int) int {
		// when arrived at finish point
		// i == m - 1 && j == n - 1
		if i == m-1 && j == n-1 {
			return 1
		}
		// 超出索引范围，return
		if i > m-1 || j > n-1 {
			return 0
		}
		return counter(i+1, j) + counter(i, j+1)
	}
	return counter(0, 0)
}

// 62. Unique Paths
// 62. 不同路径
// 思路：memoization
// time O(n) space O(n)
func uniquePathsM(m int, n int) int {
	memo := genmemo(m, n, -1)

	// counter: the number of paths of
	// from [i,j] to [m-1, n-1]
	var counter func(i, j int) int
	counter = func(i, j int) int {
		// when arrived at finish point
		// i == m - 1 && j == n - 1.
		if i == m-1 && j == n-1 {
			return 1
		}
		if i > m-1 || j > n-1 {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		memo[i][j] = counter(i+1, j) + counter(i, j+1)
		return memo[i][j]
	}

	return counter(0, 0)
}

func genmemo(m, n, initVal int) [][]int {
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}

// 62. Unique Paths
// 62. 不同路径
// 思路：dp
// time O(mn) space O(mn)
func uniquePaths(m int, n int) int {
	// dp[i][j]:represents the number of paths from [0,0] to [i,j]
	dp := gendp(m, n)

	// base case: i == 0 || j == 0, dp[i][j] = 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func gendp(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return dp
}

// 62. Unique Paths
// 62. 不同路径
// 思路：dp
// time O(mn) space O(mn)
func uniquePathsDp2(m int, n int) int {
	// dp[i][j]:represents the number of paths from [m-1,n-1] to [i,j]
	dp := gendp(m, n)

	// base case: i == m-1 || j == n-1, dp[i][j] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i+1][j] + dp[i][j+1]
			}
		}
	}
	return dp[0][0]
}
