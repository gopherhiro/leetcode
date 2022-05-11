package dp

import "math"

// 714. Best Time to Buy and Sell Stock with Transaction Fee
// 714. 买卖股票的最佳时机含手续费
// 思路：dynamic programming compress state
// time O(N) space O(1)
func maxProfit714Compress(prices []int, fee int) int {
	n := len(prices)
	// base case
	// dp[-1][0], dp[-1][1] = 0, -infinity
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dpi0 = max(dpi0, dpi1+prices[i]-fee) // 每次卖出，需要付手续费
		dpi1 = max(dpi1, dpi0-prices[i])
	}
	return dpi0
}

// 714. Best Time to Buy and Sell Stock with Transaction Fee
// 714. 买卖股票的最佳时机含手续费
// 思路：dynamic programming
// time O(N) space O(2N)
func maxProfit714(prices []int, fee int) int {
	n := len(prices)
	dp := genDp(n, 2)
	// base case
	// dp[-1][0], dp[-1][1] = 0, -infinity
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i]-fee)
			//dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
			// 根据以上状态转移方程推导所得
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee) // 每次卖出，需要付手续费
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
