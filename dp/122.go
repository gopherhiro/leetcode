package dp

// 122. Best Time to Buy and Sell Stock II
// 122. 买卖股票的最佳时机 II
// 思路：交易次数不受限，累加所有上坡之和，即可获得最大利润。
// 这是一种贪心解法
// time O(N) space O(1)
func maxProfit122(prices []int) int {
	answer := 0
	for i := 1; i < len(prices); i++ {
		// 累加所有上坡之和
		if prices[i] > prices[i-1] {
			answer += prices[i] - prices[i-1]
		}
	}
	return answer
}

// 122. Best Time to Buy and Sell Stock II
// 122. 买卖股票的最佳时机 II
// 思路：dynamic programming (状态压缩)
// time O(N) space O(1)
func maxProfitIIDPs(prices []int) int {
	n := len(prices)
	// base case
	// dp[-1][0], dp[-1][1] = 0, -infinity
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, dpi0-prices[i])
	}
	return dpi0
}

// 122. Best Time to Buy and Sell Stock II
// 122. 买卖股票的最佳时机 II
// 思路：dynamic programming
// 状态转移方程：
//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
//dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
// time O(N) space O(2N)
func maxProfitIIDP(prices []int) int {
	n := len(prices)
	dp := genDp(n, 2)
	// base case
	// dp[-1][0], dp[-1][1] = 0, -infinity
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			//dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
			//dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
			// 根据以上状态转移方程推导所得
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
