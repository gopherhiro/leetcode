package binarysearch

// 704. Binary Search
// 思路：使用二分搜索框架即可。
// time O(logN), space O(1)
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			// 中间值小于target，表示target的值在右边。
			// 故而收缩左边，即去右边找。
			left = mid + 1
		} else if nums[mid] > target {
			// 中间值大于target，表示target的值在左边。
			// 故而收缩右边，即去左边找。
			right = mid - 1
		} else if nums[mid] == target {
			return mid // 找到，直接返回。
		}
	}
	return -1
}
