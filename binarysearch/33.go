package binarysearch

import "math"

// 33. Search in Rotated Sorted Array
// 33. 搜索旋转排序数组
// 思路：「旋转数组中找目标值」 转化成「有序数组中找目标值」
// 使用基础二分搜索即可。
// time(logN)
func searchRotate(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}
		// 先确定 target 在左半段还是右半段
		// 更改mid的值，构造有序搜索范围
		// 从而将「旋转数组中找目标值」 转化成了「有序数组中找目标值」
		if target >= nums[0] {
			// target 在左半段
			// 如果mid在右半段的话，则将mid改为 max
			if nums[mid] < nums[0] {
				nums[mid] = math.MaxInt32
			}

		} else {
			// target 在右半段
			// 如果mid在左半段的话，则将mid改为 min
			if nums[mid] >= nums[0] {
				nums[mid] = math.MinInt32
			}
		}
		// 经过以上操作，已经转换为有序数组找目标值
		// 故而使用基础二分搜索解决问题即可。
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// 思路：将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。
// 此时有序部分用二分法查找。
// 无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。
// 以此类推。
func searchRotateN(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		// 确认有序部分，在有序中进行二分搜索
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}
