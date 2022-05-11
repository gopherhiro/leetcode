package dp

import "math"

// 309. Best Time to Buy and Sell Stock with cool down
// 309. 最佳买卖股票时机含冷冻期
// 思路：dynamic programming compress state
// time O(N) space O(1)
func maxProfit309Compress(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	// base case
	// dp[-2][0], dp[-1][0], dp[-1][1] = 0, 0, -infinity
	dpi0, dpi1, dpi2 := 0, math.MinInt32, 0
	for i := 0; i < n; i++ {
		tmp := dpi0
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, dpi2-prices[i])
		dpi2 = tmp
	}
	return dpi0
}

// 309. Best Time to Buy and Sell Stock with cool down
// 309. 最佳买卖股票时机含冷冻期
// 思路：dynamic programming
// time O(N) space O(2N)
func maxProfit309(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	dp := genDp(n, 2)
	// base case
	// dp[-2][0], dp[-1][0], dp[-1][1] = 0, 0, -infinity
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			//dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
			// 根据以上状态转移方程推导所得
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i-2 == -1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			//dp[i][1] = max(dp[i-1][1], dp[i-2][0] - prices[i])
			// 根据以上状态转移方程推导所得
			dp[i][1] = max(dp[i-1][1], -prices[i])
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i]) // 从 i-2 状态开始转移，表示买入具有冷冻期
	}
	return dp[n-1][0]
}
