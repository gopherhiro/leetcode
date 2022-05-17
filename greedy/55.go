package greedy

// 55. Jump Game
// 55. 跳跃游戏
// 思路：贪心思想
// 步骤：
// 1、每一步从当前位置 i 能够跳跃到的位置中，选择最远的那个跳远位置。
// 2、遍历整个数组，得到最远的跳跃长度，能够超越最后一个下标，则能够到达最后一个下标。
// - 在遍历的过程中，如果发现当前位置 i，最远的跳远位置是 i 本身，则表示从当前位置永远无法到达下一个下标，
// - 也就无法达到最后一个下标。
// time O(N) space O(1)
func canJump(nums []int) bool {
	farthest := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		curCanJumpFarthest := i + nums[i] // 从当前位置 i 能够跳到的最远的位置
		farthest = max(farthest, curCanJumpFarthest)
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
