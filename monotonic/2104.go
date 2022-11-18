package monotonic

// 2104. Sum of Subarray Ranges
// 2104. 子数组范围和
// 思路：两次遍历
// time O(n^2) space O(1)
func subArrayRanges(nums []int) int64 {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		maxVal, minVal := nums[i], nums[i]
		for j := i; j < len(nums); j++ {
			maxVal = max(maxVal, nums[j])
			minVal = min(minVal, nums[j])
			res += maxVal - minVal
		}
	}
	return int64(res)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
