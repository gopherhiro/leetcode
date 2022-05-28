package stack

// 1944. Number of Visible People in a Queue
// 1944. 队列中可以看到的人数
// 策略：单调栈
// time O(N) space O(N)
func canSeePersonsCount(heights []int) []int {
	length := len(heights)
	ans := make([]int, length)
	stacks := make([]int, 0)
	for i := length - 1; i >= 0; i-- {
		for len(stacks) > 0 && stacks[len(stacks)-1] <= heights[i] {
			ans[i]++ // 统计能够看到的人数：比Ta矮，都是Ta可以看到的。
			stacks = stacks[:len(stacks)-1]
		}
		// 栈不为空
		// 统计能够看到的人数：最后一个比Ta高的人，也是Ta可以看到的。
		if len(stacks) > 0 {
			ans[i]++
		}
		stacks = append(stacks, heights[i])
	}
	return ans
}
