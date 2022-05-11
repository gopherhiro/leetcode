package dp

import "math"

// 312. Burst Balloons
// 312. 戳气球
// 思路：dynamic programming
// dp 数组定义：dp[i][j] = x 表示：
// 戳破在开区间 (i,j) 内的所有气球，可获得的最多硬币数为 x
// base case:
// i == j, dp[i][j] == 0
// abs(i-j) == 1, dp[i][j] == 0
// time O(n^3) space O(n^2)
func maxCoins(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 添加两侧的虚拟气球对应的硬币数
	coins := make([]int, n+2)
	// 如果 i - 1或 i + 1 超出了数组的边界，
	// 那么就当它是一个硬币数为 1 的气球。
	// 即 nums[-1], nums[n] 都为 1。
	coins[0], coins[n+1] = 1, 1
	for i := 1; i <= n; i++ {
		coins[i] = nums[i-1]
	}

	// base case:
	// i == j, dp[i][j] == 0
	// abs(i-j) == 1, dp[i][j] == 0
	dp := generateDp(n+2, n+2)

	// 开始状态转移
	// i 从下到上， j 从左到右
	for i := n; i >= 0; i-- {
		for j := i + 1; j < n+2; j++ {
			// 最后戳破的气球是 k, i < k < j.
			for k := i + 1; k < j; k++ {
				// 先戳破开区间 (i,k) 和 (k,j)的气球，最后戳破的气球 k。可获得硬币数为：
				sum := dp[i][k] + dp[k][j] + coins[i]*coins[k]*coins[j]
				// 择优
				dp[i][j] = max(dp[i][j], sum)
			}
		}
	}
	return dp[0][n+1]
}

// 312. Burst Balloons
// 312. 戳气球
// 思路：回溯
// time O(n!) space O(n!)
func maxCoinsB(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := math.MinInt8
	var backtrack func(nums []int, sum int)
	backtrack = func(nums []int, sum int) {
		// 结束条件
		if len(nums) == 0 {
			ans = max(ans, sum)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 递归之前 做选择
			tmp := nums[i]
			// 计算硬币
			coin := 0
			iPrev, iNext := 1, 1
			if i-1 < 0 {
				iPrev = 1
			} else {
				iPrev = nums[i-1]
			}
			if i+1 >= len(nums) {
				iNext = 1
			} else {
				iNext = nums[i+1]
			}
			coin = iPrev * nums[i] * iNext
			nums := delete(nums, i)

			// 进入下一层决策树
			backtrack(nums, sum+coin)

			// 递归之后 撤销选择
			nums = insert(nums, i, tmp)
		}
	}
	backtrack(nums, 0)
	return ans
}

func delete(nums []int, i int) []int {
	prev := nums[:i]
	last := nums[i+1:]
	return append(prev, last...)
}

func insert(nums []int, i, elem int) []int {
	prev := nums[:i]
	last := nums[i:]
	l := append([]int{elem}, last...)
	return append(prev, l...)
}
