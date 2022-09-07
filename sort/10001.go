package sort

import "math"

// 有一个无序的整数序列，求序列的中位数
// 例如给 定数组 1、5、2、9、8、0、6，中位数是 5。
// 输入:一个无序的整数序列。
// 输出:中位数。

// 思路：基于二分搜索的快排选择
// partition 函数会将 nums[p] 排到正确的位置，使得
// nums[lo..p-1] < nums[p] < nums[p+1..hi]
// time O(N) space O(1)
func FindMedian(nums []int) int {
	// 中位数的位置：数组个数的中间位置。
	k := (len(nums) - 1) / 2
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		// 在 nums[lo..hi] 中选择一个分区点
		p := partitionN(nums, lo, hi)
		if p < k {
			// 中位数在 nums[p+1..hi] 中
			lo = p + 1
		} else if p > k {
			// 中位数在 nums[lo..p-1] 中
			hi = p - 1
		} else {
			// 找到中位数
			return nums[p]
		}
	}
	return math.MinInt16
}

func partitionN(nums []int, lo, hi int) int {
	// 选择最后一个元素作为分区点
	pivot := nums[hi]
	// i 是分区点 pivot 的下标
	// 已处理区间:[lo, i-1] 都是小于 pivot 的
	// 未处理区间:[i+1, hi-1] 都是大于 pivot 的
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			swapN(nums, i, j)
			i++
		}
	}
	// 将 pivot 放到对应的位置
	// 即 pivot 左边元素较小，右边元素较大
	swapN(nums, i, hi)

	// 返回分区点下标
	return i
}

func swapN(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
