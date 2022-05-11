package dp

/**
股票系列问题有共通性，详细解析参见：https://labuladong.github.io/algo/1/12/
1、股票系列问题状态定义：
dp[i][k][0 or 1]
0 <= i <= n - 1, 1 <= k <= K
n 为天数，大 K 为交易数的上限，0 和 1 代表是否持有股票。

2、股票系列问题通用状态转移方程和 base case：
base case：
dp[-1][...][0] = dp[...][0][0] = 0
dp[-1][...][1] = dp[...][0][1] = -infinity

3、状态转移方程：记忆方法，根据今天有没有持有股票，是选择「买入」还是「卖出」来推算帮助记忆。
dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i])
dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0] - prices[i])

特化到 k 无限制的情况，状态转移方程如下：
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
*/

// 买卖股票的最佳时机 | 通用解题模版
// 给定一个股票价格数组 prices，最多进行 t 次交易，每次交易需要支付 fee 手续费，
// 并且每次交易之后需要经过 cool 天的冷冻期才能进行下一次交易，返回你能获得最大利润。
// Find the maximum profit you can achieve.
func AllMaxProfit(prices []int, t, fee, cool int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	s := 2 // 两个状态：0（未持有股票） 和 1（持有股票）
	dp := gen3Dp(n, t+1, s)
	// base case：
	// dp[-1][...][0] = dp[...][0][0] = 0
	// dp[-1][...][1] = dp[...][0][1] = -infinity
	for i := 0; i < n; i++ {
		for k := t; k >= 1; k-- {
			// 未开始交易 base case
			if i-1 == -1 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[i]
				continue
			}

			// 冷冻期 base case
			if i-cool-1 < 0 {
				dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i]-fee) // 每次卖出，需要付手续费
				dp[i][k][1] = max(dp[i-1][k][1], -prices[i])                  // 每次买入，具有冷冻期 cool 天
				continue
			}

			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i]-fee)    // 每次卖出，需要付手续费
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-cool-1][k-1][0]-prices[i]) // 每次买入，具有冷冻期 cool 天
		}
	}
	// 总共穷举 n * t * s 个状态
	return dp[n-1][t][0]
}

// 188. Best Time to Buy and Sell Stock IV
// 188. 买卖股票的最佳时机 IV
// 思路：dynamic programming
// time O(NK) space O(2NK)
func maxProfit188(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	K, S := k, 2
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
