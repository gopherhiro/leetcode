package dp

// 416. Partition Equal Subset Sum
// 416. 分割等和子集
// 思路：dynamic programming
// dp 数组定义：dp[i][j] = x 表示：
// 对于前 i 个物品，当前背包的容量为 j 时，
// 若 x 为 true，则说明可以恰好将背包装满。
// 若 x 为 false，则说明不能恰好将背包装满。
// time O(n*sum) space O(n*sum)
func canPartition(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// 和为奇数，不能划分为两个相等的子集。
	if sum%2 != 0 {
		return false
	}
	m := sum / 2 // 子集的容量
	dp := generateBoolDp(n+1, m+1)
	// base case
	// dp[...][0] = true 背包容量为0，表示装满了。
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}
	// dp[0][...] = false 没有可选物品，则不能装满背包。
	// 创建 dp 已默认是 false
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// 默认容量不足，选择不装入。
			dp[i][j] = dp[i-1][j]
			// 容量充足，则在「装入」和「不装入」择优。
			if j-nums[i-1] >= 0 {
				// 其中一个可装满，则dp[i][j]可装满。
				unload := dp[i-1][j]
				load := dp[i-1][j-nums[i-1]]
				dp[i][j] = unload || load
			}
		}
	}
	return dp[n][m]
}

func generateBoolDp(row, col int) [][]bool {
	dp := make([][]bool, row)
	for i, _ := range dp {
		dp[i] = make([]bool, col)
	}
	return dp
}

// 416. Partition Equal Subset Sum
// 416. 分割等和子集
// 思路：recursion + memo
// dp 函数定义：dp(i,j)：
// 对于前 i 个物品，当前背包的容量为 j 时，
// 若 dp(i,j) 为 true，则说明可以恰好将背包装满。
// 若dp(i,j) 为 false，则说明不能恰好将背包装满。
// time O(n*sum) space O(n*sum)
func canPartitionM(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// 和为奇数，不能划分为两个相等的子集。
	if sum%2 != 0 {
		return false
	}
	m := sum / 2 // 子集的容量

	memo := make(map[string]bool, 0)

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// base case
		// 背包容量为 0 表示装满了。
		if j == 0 {
			return true
		}
		// 没有物品可选择，一定是装不满的
		if i == 0 {
			return false
		}

		key := fmt.Sprintf("%d,%d", i, j)
		if val, ok := memo[key]; ok {
			return val
		}

		var ret bool
		// 背包容量不够，只能选择「不装入」
		if j-nums[i-1] < 0 {
			ret = dp(i-1, j)
		}
		// 择优：装入 和 不装入
		ret = dp(i-1, j) || dp(i-1, j-nums[i-1])

		memo[key] = ret

		return ret
	}
	return dp(n, m)
}

// 416. Partition Equal Subset Sum
// 416. 分割等和子集
// 思路：recursion
// dp 函数定义：dp(i,j)：
// 对于前 i 个物品，当前背包的容量为 j 时，
// 若 dp(i,j) 为 true，则说明可以恰好将背包装满。
// 若dp(i,j) 为 false，则说明不能恰好将背包装满。
// time O(n*sum) space O(n*sum)
func canPartitionR(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return false
	}
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	// 和为奇数，不能划分为两个相等的子集。
	if sum%2 != 0 {
		return false
	}
	m := sum / 2 // 子集的容量

	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// base case
		// 背包容量为 0 表示装满了。
		if j == 0 {
			return true
		}
		// 没有物品可选择，一定是装不满的
		if i == 0 {
			return false
		}
		// 背包容量不够，只能选择「不装入」
		if j-nums[i-1] < 0 {
			return dp(i-1, j)
		}
		// 择优：装入 和 不装入
		return dp(i-1, j) || dp(i-1, j-nums[i-1])
	}
	return dp(n, m)
}
