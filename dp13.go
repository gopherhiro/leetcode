package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{3, 2, 6, 5, 0, 3}
	k := 2
	r := AllMaxProfit(nums, k, 0, 0)
	fmt.Println(r)
}

/**
股票系列问题有共通性，详细解析参见：https://labuladong.gitee.io/algo/3/26/96/
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

// 121. Best Time to Buy and Sell Stock
// 121. 买卖股票的最佳时机
// 思路：dynamic programming (状态压缩)
// time O(N) space O(1)
func maxProfitI(prices []int) int {
	n := len(prices)
	// dp[-1][0] = 0
	// dp[-1][1] = -infinity
	dpi0, dpi1 := 0, math.MinInt32
	for i := 0; i < n; i++ {
		dpi0 = max(dpi0, dpi1+prices[i])
		dpi1 = max(dpi1, -prices[i])
	}
	return dpi0
}

// 121. Best Time to Buy and Sell Stock
// 121. 买卖股票的最佳时机
// 思路：dynamic programming
// dp 数组定义：
// dp[i][0] be the maximum profit that could be gained
// at the end of the i-th day with has sell stock
// target => dp[n-1][0]
// time O(N) space O(2N)
func maxProfitDp(prices []int) int {
	n := len(prices)
	dp := genDp(n, 2)
	// base case
	// dp[-1][0], dp[-1][1] = 0, -infinity
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			// dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			// dp[i][1] = max(dp[i-1][1], -prices[i])
			// 根据状态转移方程推导所得
			// new base case
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[n-1][0]
}

func genDp(row, col int) [][]int {
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	return dp
}

// 121. Best Time to Buy and Sell Stock
// 121. 买卖股票的最佳时机
// 思路：枚举最小峰谷与其未来峰顶的差值，取其中最大差值即可。
// Remember:You can only buy one time & sell one time
// time O(N) space O(1)
func maxProfitM(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		// 维护最小价格
		minPrice = min(minPrice, prices[i])

		// 计算收益
		profit := prices[i] - minPrice

		// 更新最大收益
		maxProfit = max(maxProfit, profit)
	}
	return maxProfit
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 121. Best Time to Buy and Sell Stock
// 121. 买卖股票的最佳时机
// 思路：枚举最小峰谷与其未来峰顶的差值，取其中最大差值即可。
// Remember:You can only buy one time & sell one time
// time O(N) space O(1)
func maxProfitV(prices []int) int {
	maxProfit := 0
	minValley := prices[0]
	for i := 1; i < len(prices); i++ {
		// 维护最小峰谷
		if prices[i] < minValley {
			minValley = prices[i]
		}
		// 计算峰顶与最小峰谷差值，即收益
		profit := prices[i] - minValley
		// 更新最大收益
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// 121. Best Time to Buy and Sell Stock
// 121. 买卖股票的最佳时机
// 思路：two pointer,
// left -> buy stock
// right -> sell stock
// time O(N) space O(1)
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	answer := 0
	// left -> buy, right -> sell
	left, right := 0, 1
	for right < n {
		profit := prices[right] - prices[left]
		if profit > 0 {
			// which means we will get profit
			// so we will update our max_profit and move our right pointer alone
			answer = max(answer, profit)
			right++
		} else {
			// no profit,
			// move left pointer to the right position and increment our right pointer by 1
			left = right
			right++
		}
	}
	return answer
}

// 121. Best Time to Buy and Sell Stock
// 121. 买卖股票的最佳时机
// 思路：暴力枚举所有买卖的可能，从中选择最大profit
// time O(N^2) space O(1)
// 超时
func maxProfitF(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			profit := prices[j] - prices[i]
			res = max(res, profit)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
