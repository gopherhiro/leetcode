package array

import (
	"math"
	"math/rand"
)

// 215. Kth Largest Element in an Array
// 215. 数组中的第K个最大元素
// 策略：基于二分搜索的快排选择
// partition 函数会将 nums[p] 排到正确的位置，使得
// nums[lo..p-1] < nums[p] < nums[p+1..hi]
// time O(N) space O(1)
func findKthLargest(nums []int, k int) int {
	// 为了避免出现耗时的极端情况，先随机打乱
	shuffle(nums)
	// 排好序后，第 len(nums) - k 个，即是第K个大的元素。
	// 则转化成查找「排名第 k 的元素」
	k = len(nums) - k
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		// 在 nums[lo..hi] 中选择一个分区点
		p := partition(nums, lo, hi)
		if p < k {
			// 第 k 大元素在 nums[p+1..hi] 中
			lo = p + 1
		} else if p > k {
			// 第 k 大元素在 nums[lo..p-1] 中
			hi = p - 1
		} else {
			// 找到第 k 大元素
			return nums[p]
		}
	}
	return math.MinInt16
}

func partition(nums []int, lo, hi int) int {
	// 选择最后一个元素作为分区点
	pivot := nums[hi]
	// i 是分区点 pivot 的下标
	// 已处理区间:[lo, i-1] 都是小于 pivot 的
	// 未处理区间:[i+1, hi-1] 都是大于 pivot 的
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			swap(nums, i, j)
			i++
		}
	}
	// 将 pivot 放到对应的位置
	// 即 pivot 左边元素较小，右边元素较大
	swap(nums, i, hi)

	// 返回分区点下标
	return i
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func shuffle(nums []int) {
	rand.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})
}

// 216. Kth Smallest Element in an Array
// 216. 数组中的第K个最小元素
// 策略：基于二分搜索的快排选择
// partition 函数会将 nums[p] 排到正确的位置，使得
// nums[lo..p-1] < nums[p] < nums[p+1..hi]
// time O(N) space O(logN)
func findKthSmallest(nums []int, k int) int {
	// 为了避免出现耗时的极端情况，先随机打乱
	shuffle(nums)
	// 排好序后，找第 k-1 个元素即可。
	k = k - 1
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		// 在 nums[lo..hi] 中选择一个分区点
		p := partition(nums, lo, hi)
		if p < k {
			// 第 k 小元素在 nums[p+1..hi] 中
			lo = p + 1
		} else if p > k {
			// 第 k 小元素在 nums[lo..p-1] 中
			hi = p - 1
		} else {
			// 找到第 k 小元素
			return nums[p]
		}
	}
	return math.MinInt16
}
