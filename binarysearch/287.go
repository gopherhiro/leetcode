package binarysearch

// 287. Find the Duplicate Number
// 287. 寻找重复数
// 思路：n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内。
// 索引是有序的，所以，可以使用二分查找。
func findDuplicateQ(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 1, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1

		// 统计小于等于mid的个数
		count := 0
		for _, v := range nums {
			if v <= mid {
				count++
			}
		}
		// 如果count <= mid,表示左侧数据不存在重复的数
		if count <= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}

	}
	return left
}
