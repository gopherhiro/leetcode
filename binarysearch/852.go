package binarysearch

// 852. Peak Index in a Mountain Array
// 与 162 相同。
func peakIndexInMountainArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
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
