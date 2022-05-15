package twopointer

// 26. Remove Duplicates from Sorted Array
// 26. 删除有序数组中的重复项
// 思路：双指针
// time O(N) space O(1)
func removeDuplicate(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			// 维护 nums[0..slow] 无重复
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}
