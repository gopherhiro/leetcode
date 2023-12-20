package stack

// 402. 移掉 K 位数字
// 思路：单调递增栈
// time O(n) space O(N)
func removeKdigits(num string, k int) string {

	if len(num) == k {
		return "0"
	}

	stack := make([]rune, 0, len(num))
	for _, c := range num {

		// 维护单调递增栈
		// 只要 k > 0 且当前的 c 比栈顶的小，则栈顶出栈，k--
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > c {
			stack = stack[:len(stack)-1]
			k--
		}

		// 防止 0 入空栈，出现前导 0 的情况
		if c == '0' && len(stack) == 0 {
			continue
		}

		stack = append(stack, c)
	}

	// k 可能还大于 0，继续从 stack 删除，直到 k = 0
	for k > 0 && len(stack) > 0 {
		stack = stack[:len(stack)-1]
		k--
	}

	// 如果栈空了，直接返回 "0"
	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
