## 动态规划

- 动态规划有自底向上和自顶向下两种解决问题的方式。
- 自顶向下即记忆化递归，自底向上就是递推。

## 题目列表

- [10. 正则表达式匹配](10.go)
- [53. 最大子数组和](53.go)
- [64. 最小路径和](64.go)
- [72. 编辑距离](72.go)
- [121. 买卖股票的最佳时机](121.go)
- [122. 买卖股票的最佳时机 II](122.go)
- [123. 买卖股票的最佳时机 III](123.go)
- [174. 地下城游戏](174.go)
- [188. 买卖股票的最佳时机 IV](188.go)
- [198. 打家劫舍](198.go)
- [213. 打家劫舍 II](213.go)
- [300. 最长递增子序列](300.go)
- [309. 最佳买卖股票时机含冷冻期](309.go)
- [312. 戳气球](312.go)
- [322. 零钱兑换](322.go)
- [354. 俄罗斯套娃信封问题](354.go)
- [416. 分割等和子集](416.go)
- [509. 斐波那契数](509.go)
- [516. 最长回文子序列](516.go)
- [518. 零钱兑换 II](518.go)
- [583. 两个字符串的删除操作](583.go)
- [651. 4键键盘](651.go)
- [712. 两个字符串的最小ASCII删除和](712.go)
- [714. 买卖股票的最佳时机含手续费](714.go)
- [ 887. 鸡蛋掉落](887.go)
- [931. 下降路径最小和](931.go)
- [1143. 最长公共子序列](1143.go)
- [1312. 让字符串成为回文串的最少插入次数](1312.go)

## 股票系列问题解题攻略

```go

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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

```