package greedy

// 45. Jump Game II
// 45. 跳跃游戏 II
// 策略：贪心
// 思路：每一次都选择跳得最远的位置，作为起跳点，即可保证最小的跳跃次数。
// time O(N) space O(1)
func jump(nums []int) int {
	jump, farthest, end := 0, 0, 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end {
			jump++         // 起跳
			end = farthest // 标记本次可以跳到的最远位置
		}
	}
	return jump
}
