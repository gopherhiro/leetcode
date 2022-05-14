package sort

import "math/rand"

// 912. Sort an Array
// 912. 排序数组
// 策略：快排
// time O(N*logN) space O(logN)
func quickSort(nums []int) []int {
	n := len(nums)
	if n == 0 || n == 1 {
		return nums
	}
	var sort func(num []int, lo, hi int)
	sort = func(nums []int, lo, hi int) {
		// base case
		if lo >= hi {
			return
		}
		p := partition(nums, lo, hi)
		sort(nums, lo, p-1)
		sort(nums, p+1, hi)
	}
	// 为了避免出现耗时的极端情况，先随机打乱
	shuffle(nums)

	sort(nums, 0, n-1)
	return nums
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
