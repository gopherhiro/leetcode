package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{5, 0, 0, 0}
	fmt.Println(res)
}

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

// 1732. 找到最高海拔
// time O(N) space O(1)
func largestAltitude(gain []int) int {
	largest := 0
	preSum := make([]int, len(gain)+1)
	for i := 1; i <= len(gain); i++ {
		preSum[i] = preSum[i-1] + gain[i-1]
		largest = max(largest, preSum[i])
	}

	return largest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 1413. Minimum Value to Get Positive Step by Step Sum
// 1413. 逐步求和得到正数的最小值
// time O(N) space O(1)
// Algorithm:
// 1.Traverse the array nums and calculate every step-by-step total,
// use total to record the current step-by-step total,
// and minVal to record the minimum step-by-step total.
// 2.Return -minVal + 1, that is the minimum valid startValue.
func minStartValue(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	minVal := 0
	total := 0
	for _, num := range nums {
		total += num
		minVal = min(minVal, total)
	}

	return 1 - minVal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 724. Find Pivot Index
// 724. 寻找数组的中心下标
// time O(N) space O(N)
func pivotIndexI(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	total := 0
	for _, num := range nums {
		total += num
	}

	// leftSum + num[i] + rightSum == total
	// if leftSum == rightSum
	// so that 2 * sum + num[i] == total
	// so that i is pivot index
	sum := 0 // sum of left
	for i := 0; i < len(nums); i++ {
		if 2*sum+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}

// 724. Find Pivot Index
// 724. 寻找数组的中心下标
// 思路：数组中心下标：左侧所有元素相加的和等于右侧所有元素相加的和。
// 遍历数组，利用前缀和求出下标左右侧元素之和，比较左右两侧元素和是否相等。
// 若相等，则找到中心下标。
// time O(N) space O(N)
// 注意：
// preSum[i] 对应 num[0..i-1]的和。
// num[i..j] 的和为 preSum[j+1] - preSum[i]
func pivotIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	// 构建前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 遍历数组，找寻中心下标记
	n := len(nums)
	for i := 0; i < n; i++ {
		leftSum := preSum[i] - preSum[0]    // nums[0..i-1] 的和
		rightSum := preSum[n] - preSum[i+1] // nums[i+1..n-1] 的和
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

// 303. Range Sum Query - Immutable
// 303. 区域和检索 - 数组不可变
// time O(N), space O(N)
type NumArray struct {
	preSum []int
}

// 构造前缀和
func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	return NumArray{preSum: preSum}
}

// 计算索引 [left, right] 之间的 nums 元素的和
func (na *NumArray) SumRange(left int, right int) int {
	return na.preSum[right+1] - na.preSum[left]
}

// 560. Subarray Sum Equals K
// 560. 和为 K 的子数组
// 思路：前缀和+hash
// time O(N) space O(N)
func subarraySum1(nums []int, k int) int {
	hash := make(map[int]int, 0)
	hash[0] = 1
	ans, sum0i := 0, 0
	for i := 0; i < len(nums); i++ {
		sum0i += nums[i]   // sum[0..i]
		sum0j := sum0i - k // sum[0..j]
		if count, ok := hash[sum0j]; ok {
			ans += count
		}
		// sum[0..i]加入hash，并记录出现的次数。
		hash[sum0i] += 1
	}
	return ans
}

// 560. Subarray Sum Equals K
// 560. 和为 K 的子数组
// 思路：前缀和+穷举
// time O(N^2) space O(N)
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	// 构造前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 穷举所有子数组
	ans := 0
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			// nums[j..i-1]
			if preSum[i]-preSum[j] == k {
				ans++
			}
		}
	}

	return ans
}
