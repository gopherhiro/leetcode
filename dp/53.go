package dp

// 53. Maximum Subarray
// 53. 最大子数组和
// 思路：动态规划(状态压缩)
// time O(N) space O(1)
// dp[i]定义：以nums[i]结尾的最大子数组和为dp[i]。
// dp[i] 有两种「选择」：
// 1、与前面的相邻子数组连接，形成一个和更大的子数组；
// 2、不与前面的子数组连接，自己作为一个子数组。
// 在这两种选择中择优，就可以计算出最大子数组和。
// 备注：因为 dp[i] 只与 dp[i-1]有关，
// 故进行状态压缩(使用变量替换数组进行迭代)，可减少空间复杂度
func maxSubArrayO(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp0 := nums[0] // base case
	answer := dp0
	for i := 1; i < len(nums); i++ {
		dp1 := max(nums[i], nums[i]+dp0) // 状态转移方程
		answer = max(answer, dp1)        // 选取子数组和 的最大
		dp0 = dp1                        // 变量迭代
	}
	return answer
}

// 53. Maximum Subarray
// 53. 最大子数组和
// 思路：动态规划
// time O(N) space O(N)
// dp[i]定义：以nums[i]结尾的最大子数组和为dp[i]。
// dp[i] 有两种「选择」：
// 1、与前面的相邻子数组连接，形成一个和更大的子数组；
// 2、不与前面的子数组连接，自己作为一个子数组。
// 在这两种选择中择优，就可以计算出最大子数组和。
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))

	// base case
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		// 状态转移方程
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	// 找寻最大子数组的和
	answer := dp[0]
	for i := 1; i < len(dp); i++ {
		answer = max(answer, dp[i])
	}
	return answer
}

// 53. Maximum Subarray
// 53. 最大子数组和
// 思路：Recursion
// time O(N) space O(N)
func maxSubArrayR(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	ans := nums[0] // 默认第一个元素即是结果

	var helper func(i int) int
	helper = func(i int) int {
		// base case
		if i < 0 {
			return 0
		}
		if i == 0 {
			return nums[0]
		}
		// opt 有两种选择：选择值 max 的即可。
		// 1、与前面的相邻子数组连接，形成一个和更大的子数组。
		// 2、自己作为一个子数组。
		opt := max(nums[i], nums[i]+helper(i-1))
		ans = max(ans, opt)
		return opt
	}
	helper(len(nums) - 1)
	return ans
}
