package twopointer

// 167. Two Sum II - Input Array Is Sorted
// 167. 两数之和 II - 输入有序数组
// 思路：双指针（左右指针）。
// time O(N) space O(1)
func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left, right}
		} else if sum < target {
			left++
		} else if sum > target {
			right--
		}
	}
	return []int{-1, -1}
}
