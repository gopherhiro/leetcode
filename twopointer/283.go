package twopointer

// 283. Move Zeroes
// 283. 移动零
// 思路：two pointer
// time O(N) space O(1)
func moveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		// 使用 slow 指针标记值为 0 的元素位置，
		// 使用 fast 指针找到值不为 0 的元素的位置
		// 交换两个即可。
		if nums[fast] != 0 {
			// [0..slow]都不为0
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}
