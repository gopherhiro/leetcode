package binarysearch

// 有一个已经排好序的整数序列(升序，无重复项)，序列中可能有正整数、负整数或者0，请用你认为 最优的方法求序列中绝对值最小的数。
// 输入: 一个有序的整数序列。
// 输出: 绝对值最小的数。
// 升序序列，可以考虑二分搜索。
// time O(logN) space O(1)
func minAbs(nums []int) int {
	n := len(nums)
	// 只有一个元素时
	if n == 1 {
		return nums[0]
	}
	// 都是正数
	if nums[0] > 0 {
		return nums[0]
	}
	// 都是负数
	if nums[n-1] <= 0 {
		return nums[n-1]
	}

	// 有正有负时，二分搜索
	target := 0
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return nums[mid]
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	if abs(nums[left]) < abs(nums[right]) {
		return nums[left]
	}
	return nums[right]
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
