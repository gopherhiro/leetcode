package dp

import "math"

// 322. Coin Change
// 322. 零钱兑换
// 思路：迭代
// time O(kN) space O(N)
// 备注：暂时没理解全，先放着
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for k := 0; k < len(dp); k++ {
		dp[k] = amount + 1
	}
	dp[0] = 0
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 322. Coin Change
// 322. 零钱兑换
// 思路：迭代
// time O(kN) space O(N)
func coinChangeL(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for k := 0; k < len(dp); k++ {
		dp[k] = amount + 1
	}
	dp[0] = 0 // base case
	for i := 0; i < len(dp); i++ {
		// 使用数学归纳法来理解就容易得多啦
		// 假设 dp[0...i-1] 都已经被算出来了，然后问自己：怎么通过这些结果算出 dp[i]？
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 322. Coin Change
// 322. 零钱兑换
// 思路：带备忘录的递归（自顶向下）
// time O(kN) space O(N)
func coinChangeM(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}

	const NoSolution = -1
	memo := make(map[int]int, 0)
	var helper func(amount int) int
	helper = func(amount int) int {
		// base case
		if amount < 0 {
			return NoSolution
		}
		if amount == 0 {
			return 0
		}
		// 先查询备忘录，不存在再计算。
		// 备忘录有，则直接返回。
		if _, ok := memo[amount]; ok {
			return memo[amount]
		}

		// 遍历所有可能的状态选择，求最小值
		minCount := math.MaxInt32
		for _, coin := range coins {
			subCount := helper(amount - coin)
			// 跳过无解子问题
			if subCount == NoSolution {
				continue
			}
			minCount = min(minCount, subCount+1)
		}
		// [最少的硬币个数]存储到备忘录中
		memo[amount] = minCount
		if minCount == math.MaxInt32 {
			memo[amount] = NoSolution
		}
		return memo[amount]
	}

	return helper(amount)
}

// 322. Coin Change
// 322. 零钱兑换
// 思路：暴力递归（超时）
// time O(k*3^N) space O(3^N)
func coinChangeR(coins []int, amount int) int {
	// base case
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	// 遍历所有可能的状态选择，求最小值
	minCount := math.MaxInt32
	for _, coin := range coins {
		subCount := coinChangeR(coins, amount-coin)
		// 跳过无解子问题
		if subCount == -1 {
			continue
		}
		minCount = min(minCount, subCount+1)
	}
	if minCount == math.MaxInt32 {
		return -1
	}
	return minCount
}
