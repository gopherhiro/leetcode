package dp

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
