package dp

import "math"

// 64. Minimum Path Sum
// 64. 最小路径和
// 思路：recursion
// dp 函数定义：从左上角位置 (0, 0) 走到位置 (i, j) 的最小路径和为dp(i,j)
// time O(m*n) space O(m*n)
func minPathSumM(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	memo := genmemo(m, n, -1)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		// 如果索引出界，返回一个很大的值，
		// 保证在取 min 的时候不会被取到
		if i < 0 || j < 0 {
			return math.MaxInt32
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// 选取左边、上面的最小路径和的最小值
		// 加上 grid[i][j]，就是到达 (i, j) 的最小路径和。
		up := dp(i-1, j)
		left := dp(i, j-1)
		ret := min(up, left) + grid[i][j]
		memo[i][j] = ret
		return ret
	}
	return dp(m-1, n-1)
}

func genmemo(m, n, initVal int) [][]int {
	memo := make([][]int, m)
	for i, _ := range memo {
		memo[i] = make([]int, n)
		for j, _ := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}

// 64. Minimum Path Sum
// 64. 最小路径和
// 思路：recursion
// dp 函数定义：从左上角位置 (0, 0) 走到位置 (i, j) 的最小路径和为dp(i,j)
// time O(m*n) space O(m*n)
func minPathSumR(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == 0 && j == 0 {
			return grid[0][0]
		}
		// 如果索引出界，返回一个很大的值，
		// 保证在取 min 的时候不会被取到
		if i < 0 || j < 0 {
			return math.MaxInt32
		}
		// 选取左边、上面的最小路径和的最小值
		// 加上 grid[i][j]，就是到达 (i, j) 的最小路径和。
		up := dp(i-1, j)
		left := dp(i, j-1)
		return min(up, left) + grid[i][j]
	}
	return dp(m-1, n-1)
}

// 64. Minimum Path Sum
// 64. 最小路径和
// 思路：dynamic programming
// dp 数组定义：从左上角位置 (0, 0) 走到位置 (i, j) 的最小路径和为dp[i][j]
// time O(m*n) space O(m*n)
// 备注：自己独立写出来的第一个动态规划题。
func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	dp := generateDp(m, n)
	// base case
	// dp[0][...]
	dp[0][0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	// dp[...][0]
	for i := 1; i < m; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	// 状态转移
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			up := dp[i-1][j]
			left := dp[i][j-1]
			dp[i][j] = min(up, left) + grid[i][j]
		}
	}
	return dp[m-1][n-1]
}
