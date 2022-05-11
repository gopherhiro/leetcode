package dp

import "math"

// 322. Coin Change
// 322. 零钱兑换
// 思路：迭代
// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出。
// 则要求的目标为：dp[amount]
// time O(kN) space O(N)
func coinChangeL(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// init dp 值为 amount + 1
	// 凑成 amount 金额的硬币数最多只可能等于 amount，
	// 所以初始化为 amount + 1 就相当于初始化为正无穷，便于后续取最小值。
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
	// 根据 dp 数组的定义，要求的目标为 dp[amount]
	if dp[amount] == amount+1 {
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

	memo := make(map[int]int, 0)
	var helper func(amount int) int
	helper = func(amount int) int {
		// base case
		if amount < 0 {
			return -1
		}
		if amount == 0 {
			return 0
		}
		// 备忘录有，则直接返回。
		if count, ok := memo[amount]; ok {
			return count
		}

		// 遍历所有可能的状态选择，求最小值
		minCount := math.MaxInt32
		for _, coin := range coins {
			subCount := helper(amount - coin)
			// 跳过无解子问题
			if subCount == -1 {
				continue
			}
			minCount = min(minCount, subCount+1)
		}
		// 存储到备忘录中
		if minCount == math.MaxInt32 {
			memo[amount] = -1
		} else {
			memo[amount] = minCount
		}

		return memo[amount]
	}

	return helper(amount)
}

// 322. Coin Change
// 322. 零钱兑换
// 思路：暴力递归（超时）
// dp函数定义：要凑出金额 amount，至少要 dp(amount) 个硬币
// time O(k^N) space O(k^N)
func coinChangeR(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	var dp func(amount int) int
	dp = func(amount int) int {
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
			// 计算子问题的结果
			subCount := dp(amount - coin)
			// 跳过无解子问题
			if subCount == -1 {
				continue
			}
			// 子问题结果 +1，择优
			minCount = min(minCount, subCount+1)
		}
		if minCount == math.MaxInt32 {
			return -1
		}
		return minCount
	}
	// 要求的结果是：dp(amount)
	return dp(amount)
}
