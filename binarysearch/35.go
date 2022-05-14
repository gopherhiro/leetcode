package binarysearch

// 35. 搜索插入位置
// 35. Search Insert Position
// 思路：搜索左侧边界的二分搜索。
// 搜索左侧边界的二分搜索，当目标元素 target 不存在数组 nums 中时，
// 搜索左侧边界的二分搜索的返回值可以做以下几种解读：
// 1、返回的这个值是 nums 中大于等于 target 的最小元素索引。
// 2、返回的这个值是 target 应该插入在 nums 中的索引位置。
// 3、返回的这个值是 nums 中小于 target 的元素个数。
// time O(logN) space O(1)
func searchInsert(nums []int, target int) int {
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
	return left
}
