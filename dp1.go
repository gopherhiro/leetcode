package main

import (
	"fmt"
	"math"
	"sort"
)

/****** 动态规划有自底向上和自顶向下两种解决问题的方式。******/
/****** 自顶向下即记忆化递归，自底向上就是递推。 ******/

func main() {
	nums := []int{-2, -1}
	r := maxSubArrayR(nums)
	fmt.Println(r)
}

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

// 354. Russian Doll Envelopes
// 354. 俄罗斯套娃信封问题
// 思路：最长递增子序列
// 1、根据 w 升序排序，if w 相等，则根据 h 降序排列
// 2、把所有的 h作为一个数组，寻找[最长递增子序列]即是[最大信封嵌套个数]
// time O(N*logN) space O(N)
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	// 根据 w 升序排序，if w 相等，则根据 h 降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})

	// 获取 height 数组
	height := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		height[i] = envelopes[i][1]
	}
	// 求 height longest incr subsequence
	return lengthOfLISBinary(height)
}

// 300. Longest Increasing Subsequence
// 300. 最长递增子序列
// 思路：二分搜索（原理：耐心排序）
// time O(N*logN) space O(N)
func lengthOfLISBinary(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	piles := 0                    // 堆数，即是最长递增子序列的长度
	top := make([]int, len(nums)) // 堆顶元素组成的列表
	for i := 0; i < len(nums); i++ {
		target := nums[i]

		/***** 搜索左侧边界的二分查找 *****/
		left, right := 0, piles-1
		for left <= right {
			mid := left + (right-left)/2
			if top[mid] < target {
				left = mid + 1
			} else if top[mid] > target {
				right = mid - 1
			} else if top[mid] == target {
				right = mid - 1
			}
		}
		// 没有找到合适的堆，则新建一堆
		if left == piles {
			piles++
		}
		// 把 target 放到搜索到的堆顶
		top[left] = target
	}
	// 堆数即是 LIS 长度
	return piles
}

// 300. Longest Increasing Subsequence
// 300. 最长递增子序列
// 思路：动态规划
// time O(N^2) space O(N)
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 定义 dp[i]：以 nums[i] 为结尾的最长递增子序列的长度。
	// 则 dp 数组中最大的那个值就是最长的递增子序列长度。
	// 初始化 base case：dp[i] 至少为 1。
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	// 利用数学归纳法思想来思考，会更容易理解。
	// 假设 dp[0...i-1] 都已经被算出来了，然后怎么通过这些结果算出 dp[i]？
	// 假设 dp[0..4] 都已经知道了，如何通过这些已知结果推出 dp[5] 呢？
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	// 遍历dp找寻最大值，即得整个数组中的：最长递增子序列
	maxSubSeqLen := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > maxSubSeqLen {
			maxSubSeqLen = dp[i]
		}
	}
	return maxSubSeqLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：迭代（自下而上，动态压缩）
// time O(N) space O(1)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	// f(i) 只与 f(i-1)、f(i-2)两种状态相关
	// 故而进行状态压缩，迭代两个状态即可。
	// 从而大大降低空间复杂度。
	prev, curr := 0, 1
	for i := 2; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：动态迭代（自下而上）
// time O(N) space O(N)
func fibDP(n int) int {
	if n <= 1 {
		return n
	}
	// 自下而上，逐步迭代
	dp := make(map[int]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：带备忘录的递归（自顶而下）
// time O(N) space O(N)
func fibM(n int) int {
	memo := make(map[int]int, n+1)
	var helper func(n int) int
	helper = func(n int) int {
		if n <= 1 {
			return n
		}
		// 先查询备忘录，存在直接返回。
		if _, ok := memo[n]; ok {
			return memo[n]
		}
		// 计算并存储到备忘录中，并返回值
		memo[n] = helper(n-1) + helper(n-2)
		return memo[n]
	}

	return helper(n)
}

// 509. Fibonacci Number
// 509. 斐波那契数
// 思路：暴力递归
// time O(2^N) space O(1)
func fibR(n int) int {
	if n <= 1 {
		return n
	}
	return fibR(n-1) + fibR(n-2)
}
