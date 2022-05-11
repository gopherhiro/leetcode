package dp

import "math"

// 123. Best Time to Buy and Sell Stock III
// 123. 买卖股票的最佳时机 III
// 思路：dynamic programming compress state
// time O(N) space O(1)
func maxProfit123compress(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	// base case：
	// dp[-1][...][0] = dp[...][0][0] = 0
	// dp[-1][...][1] = dp[...][0][1] = -infinity
	dpi10, dpi11 := 0, math.MinInt32
	dpi20, dpi21 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dpi10 = max(dpi10, dpi11+prices[i])
		dpi11 = max(dpi11, -prices[i])
		dpi20 = max(dpi20, dpi21+prices[i])
		dpi21 = max(dpi21, dpi10-prices[i])
	}
	return dpi20
}

// 123. Best Time to Buy and Sell Stock III
// 123. 买卖股票的最佳时机 III
// 思路：dynamic programming
// time O(NK) space O(2NK)
func maxProfit123(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	K, S := 2, 2
	dp := gen3Dp(n, K+1, S)
	// base case：
	// dp[-1][...][0] = dp[...][0][0] = 0
	// dp[-1][...][1] = dp[...][0][1] = -infinity
	for i := 0; i < n; i++ {
		for k := K; k >= 1; k-- {
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	// 总共穷举 n * K * S 个状态
	return dp[n-1][K][0]
}

func gen3Dp(x, y, z int) [][][]int {
	dp := make([][][]int, x)
	for i := range dp {
		dp[i] = make([][]int, y)
		for j := range dp[i] {
			dp[i][j] = make([]int, z)
		}
	}
	return dp
}
