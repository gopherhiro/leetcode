package stack

// 739. Daily Temperatures
// 739. 每日温度
// 策略：单调栈
// time O(N) space O(N)
func dailyTemperatures(tpr []int) []int {
	length := len(tpr)
	ans := make([]int, length) // 记录结果
	stacks := make([]int, 0)   // 记录元素索引
	for i := length - 1; i >= 0; i-- {
		// 保证栈单调特性
		for len(stacks) > 0 && tpr[stacks[len(stacks)-1]] <= tpr[i] {
			stacks = stacks[:len(stacks)-1]
		}

		// 栈不空
		// 栈顶元素与当前元素之差，即索引之差。
		// 表示从当前日期，需要经过 top - i 天后，会出现更高温度。
		if len(stacks) > 0 {
			top := stacks[len(stacks)-1]
			ans[i] = top - i
		}
		// 把索引存储到栈中
		stacks = append(stacks, i)
	}
	return ans
}
