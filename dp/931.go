package dp

import "math"

// 931. Minimum Falling Path Sum
// 931. 下降路径最小和
// 思路：dynamic programming
// dp 数组定义：从第一行下降，下降到 matrix[i][j] 的最小路径和为 dp[i][j]
// time O(n^2) space O(n^2)
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	dp := generateDp(n, n)
	// base case : dp[0][...] = matrix[0][j]
	for j := 0; j < n; j++ {
		dp[0][j] = matrix[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			// 边界越界情况处理：左边界
			// j - 1 < 0
			if j == 0 {
				upMid, upRight := dp[i-1][j], dp[i-1][j+1]
				dp[i][j] = min(upMid, upRight) + matrix[i][j]
				continue
			}
			// 边界越界情况处理：右边界
			// j + 1 >= n
			if j == n-1 {
				upLeft, upMid := dp[i-1][j-1], dp[i-1][j]
				dp[i][j] = min(upLeft, upMid) + matrix[i][j]
				continue
			}
			// 非边界情况，则正常选择
			upLeft, upMid, upRight := dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]
			dp[i][j] = min(upLeft, min(upMid, upRight)) + matrix[i][j]
		}
	}
	// 终点在最后一行，最后一行中找寻 min of dp[i][j]
	// 即是 minPathSum
	minPathSum := math.MaxInt32
	for j := 0; j < n; j++ {
		minPathSum = min(minPathSum, dp[n-1][j])
	}
	return minPathSum
}

// 931. Minimum Falling Path Sum
// 931. 下降路径最小和
// 思路：记忆化递归
// dp 函数定义：从第一行下降，下降到 matrix[i][j] 的最小路径和为 dp(i,j)
// time O(n^2) space O(n^2)
func minFallingPathSumM(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}

	memo := genmemo(n, n, math.MinInt32)

	var dp func(i, j int) int
	dp = func(i, j int) int {
		// 数组越界，返回一个不可能被取到的最大值即可。
		if i < 0 || j < 0 || i >= n || j >= n {
			return math.MaxInt64
		}
		// 根据 dp 函数定义，从 matrix[0][j] 开始下落，如果落到的目的地就是 i == 0
		// 则所需的路径和为 matrix[0][j]
		// base case
		if i == 0 {
			return matrix[0][j]
		}

		if memo[i][j] != math.MinInt32 {
			return memo[i][j]
		}
		upLeft, upMid, upRight := dp(i-1, j-1), dp(i-1, j), dp(i-1, j+1)
		memo[i][j] = min(upLeft, min(upMid, upRight)) + matrix[i][j]
		return memo[i][j]
	}

	// 终点在最后一行的任意一列
	minPathSum := math.MaxInt32
	for j := 0; j < n; j++ {
		minPathSum = min(minPathSum, dp(n-1, j))
	}
	return minPathSum
}

// 931. Minimum Falling Path Sum
// 931. 下降路径最小和
// 思路：暴力递归
// dp 函数定义：从第一行下降，下降到 matrix[i][j] 的最小路径和为 dp(i,j)
// time O(2^n) space O(2^n)
func minFallingPathSumR(matrix [][]int) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// 数组越界，返回一个不可能被取到的最大值即可。
		if i < 0 || j < 0 || i >= n || j >= n {
			return math.MaxInt64
		}
		// 根据 dp 函数定义，从 matrix[0][j] 开始下落，如果落到的目的地就是 i == 0
		// 则所需的路径和为 matrix[0][j]
		// base case
		if i == 0 {
			return matrix[0][j]
		}
		upLeft, upMid, upRight := dp(i-1, j-1), dp(i-1, j), dp(i-1, j+1)
		return min(upLeft, min(upMid, upRight)) + matrix[i][j]
	}

	// 终点在最后一行的任意一列
	minPathSum := math.MaxInt32
	for j := 0; j < n; j++ {
		minPathSum = min(minPathSum, dp(n-1, j))
	}
	return minPathSum
}
