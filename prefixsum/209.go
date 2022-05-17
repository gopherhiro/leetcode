package prefixsum

import "math"

// 209. Minimum Size Subarray Sum
// 209. 长度最小的子数组
// 思路：two pointer
// time O(N) space O(1)
// We could keep 2 pointer,
// one for the start and another for the end of the current subarray,
// and make optimal moves so as to keep the sum greater than target
// as well as maintain the lowest size possible.
func minSubArrayLenO(target int, nums []int) int {
	ans := math.MaxInt32
	left, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for sum >= target {
			// where i+1−left is the size of current subarray
			ans = min(ans, i-left+1)
			sum -= nums[left]
			left++ // 收缩左边界
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}

// 209. Minimum Size Subarray Sum
// 209. 长度最小的子数组
// 思路：前缀和
// time O(N^2) space O(N)
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 构造前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 穷举所有子数组
	minLength := math.MaxInt32
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] >= target {
				length := i - j
				minLength = min(minLength, length)
			}
		}
	}

	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}
