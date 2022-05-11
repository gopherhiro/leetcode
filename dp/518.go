package dp

// 518. Coin Change 2
// 518. 零钱兑换 II
// 思路：dynamic programming（状态压缩之后）
// dp 数组定义：使用coins中的前 i 个硬币，凑成总金额 j，总共有 dp[i][j] 种凑法。
// time O(n*amount) space O(amount)
func changeD(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

// 518. Coin Change 2
// 518. 零钱兑换 II
// 思路：dynamic programming
// dp 数组定义：使用coins中的前 i 个硬币，凑成总金额 j，总共有 dp[i][j] 种凑法。
// time O(n*amount) space O(n*amount)
func change(amount int, coins []int) int {
	n := len(coins)
	// 没有硬币，则无法凑出总金额。
	if n == 0 {
		return 0
	}
	// 总金额为0，只有一种方法：什么也不干，即可凑出总金额为0
	if amount == 0 {
		return 1
	}

	// 两个状态（可选择的硬币，总金额），故而需要二维dp数组。
	// target : dp[n][amount]
	dp := generateDp(n+1, amount+1)

	// base case
	// dp[0][...] = 0
	// dp[...][0] = 1
	// 总金额为0，只有一种方法：什么也不干，即可凑出总金额为0
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j] // 不装进背包
			if j-coins[i-1] >= 0 {
				unload := dp[i-1][j]        // 不装进背包
				load := dp[i][j-coins[i-1]] // 装进背包。由于对应面值的硬币数无限，所以关注如何凑出j-coins[i-1]即可。
				dp[i][j] = unload + load
			}
		}
	}
	return dp[n][amount]
}

func generateDp(row, col int) [][]int {
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	return dp
}
