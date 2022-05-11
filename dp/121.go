package dp

import "math"

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
