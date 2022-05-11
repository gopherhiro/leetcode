package dp

// 198. House Robber
// 198. 打家劫舍
// 思路：dynamic programming (状态压缩)
// dp 数组定义：表示从 start 开始能够偷窃到的最高金额为 dp[start]。
// 根据定义，dp[0] 即是我们要求的结果
// time O(N) space O(1)
func robDps(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// base case
	// dp[n],dp[n+1] := 0,0
	dpi, dpi1, dpi2 := 0, 0, 0
	for i := n - 1; i >= 0; i-- {
		dpi = max(dpi1, nums[i]+dpi2)
		dpi2 = dpi1
		dpi1 = dpi
	}
	return dpi
}

// 198. House Robber
// 198. 打家劫舍
// 思路：dynamic programming
// dp 数组定义：表示从 start 开始能够偷窃到的最高金额为 dp[start]。
// 根据定义，dp[0] 即是我们要求的结果
// time O(N) space O(N)
func robDp(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// base case
	// dp[n] = 0
	dp := make([]int, n+2)
	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], nums[i]+dp[i+2])
	}
	return dp[0]
}

// 198. House Robber
// 198. 打家劫舍
// 思路：recursion
// dp 函数定义：表示从 start 开始能够偷窃到的最高金额为 dp(start)。
// time O(2^n) space O(2^n)
func robR(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	memo := generateMemoN(n, -1)

	var dp func(start int) int
	dp = func(start int) int {
		// base case
		if start >= n {
			return 0
		}

		if memo[start] != -1 {
			return memo[start]
		}

		take := nums[start] + dp(start+2) // 取
		notTake := dp(start + 1)          // 不取
		ret := max(take, notTake)
		memo[start] = ret

		return ret
	}
	return dp(0)
}

func generateMemoN(n, initVal int) []int {
	memo := make([]int, n)
	for i := range memo {
		memo[i] = initVal
	}
	return memo
}
