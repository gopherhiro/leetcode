package twopointer

// 备注：这道题有点不好理解，可以忽略。
// 80. Remove Duplicates from Sorted Array II
// 80. 删除有序数组中的重复项 II
// 思路：two pointer
// 原地修改：
// 1、slow 指针指向当前即将放置元素的位置。
// 2、fast 指针向后遍历所有元素。
// time O(N) space O(1)
func removeDuplicatesII(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		// 因为数组是有序的，如果 num[fast] == num[slow-2]
		// 则出现 num[fast] == num[slow-2] == num[slow-1]
		// 存在连续的多于2个的重复元素。
		if nums[fast] == nums[slow-2] {
			// 所以，当 num[fast] == num[slow-2]时
			// slow 不动，标记当前需要删除元素的位置
			// fast 则继续后移，找寻新元素。
			fast++
		} else {
			// 使用新元素替换 slow 位置的元素
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}
	// slow 即是去重后的数组长度
	return slow
}
