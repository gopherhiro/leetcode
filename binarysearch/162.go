package binarysearch

// 162. Find Peak Element
// 162. 寻找峰值
// 思路：二分搜索
func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	// A peak element is an element
	// that is strictly greater than its neighbors.
	// the last element, which does not meet the requirements, is ignored.
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[mid+1] {
			// in a descending sequence of numbers. or a local falling slope
			// reduce the search space to the left of mid(including itself)
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
