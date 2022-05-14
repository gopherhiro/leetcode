package binarysearch

// 410. Split Array Largest Sum
// 410. 分割数组的最大值
// 思路：二分搜索之左侧边界，寻找子数组和最大的最小值。
// time O(N*logM)，N 是nums元素的个数，M是其所有元素之和。
// 难点在于实际问题如何转换为二分搜索问题。
func splitArray(nums []int, m int) int {
	if len(nums) == 0 || m == 0 {
		return 0
	}
	// 子数组至少包含一个元素，至多包含整个数组
	// 题目要求：子数组和最大，所以：
	// 「最大」子数组和的取值范围就是闭区间[max(nums), sum(nums)]
	left, right := getMax(nums), getSum(nums)
	for left <= right {
		mid := left + (right-left)/2
		count := splitCount(nums, mid)
		if count < m {
			// 最大子数组和上限高了，减小一些，所以，收缩右边界。
			right = mid - 1
		} else if count > m {
			// 最大子数组和上限低了，增加一些，所以，收缩左边界。
			left = mid + 1
		} else if count == m {
			// 收缩右边界，从而锁定左边界。
			right = mid - 1
		}
	}
	return left
}

// 若限制最大子数组和为 max，
// 计算 nums 至少可以被分割成几个子数组
func splitCount(nums []int, max int) int {
	sum, count := 0, 1
	for _, v := range nums {
		if sum+v > max {
			count++
			sum = 0
		}
		sum += v
	}
	return count
}
