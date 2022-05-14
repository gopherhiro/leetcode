package binarysearch

// 34. Find First and Last Position of Element in Sorted Array
// 34. 在排序数组中查找元素的第一个和最后一个位置
// 思路：使用二分搜索框架
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := leftBound(nums, target)
	right := rightBound(nums, target)
	return []int{left, right}
}

// 搜索左侧边界的二分搜索
func leftBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 不返回，收缩右边界，从而锁定左边界。
			right = mid - 1
		}
	}
	// 检查越界情况
	if left >= len(nums) || nums[left] != target {
		return -1
	}
	return left
}

// 搜索右侧边界的二分搜索
func rightBound(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			// 不返回，收缩左边界，从而锁定右边界。
			left = mid + 1
		}
	}
	// 检查越界情况
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}
