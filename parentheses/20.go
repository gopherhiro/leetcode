package parentheses

// 20. Valid Parentheses
// 20. 有效的括号
// time O(N), space O(N)
func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]byte, 0)
	// 遍历字符串
	for i := 0; i < len(s); i++ {
		// 是左括号，入栈
		if isLeft(s[i]) {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}

		// 遇见右括号，出栈
		top := stack[len(stack)-1]
		if leftOf(s[i]) == top {
			stack = stack[:len(stack)-1]
		} else {
			// 不匹配，括号无效。
			return false
		}
	}
	// 遍历完 s，栈空，括号有效。
	if len(stack) == 0 {
		return true
	}
	return false
}

func leftOf(ch byte) byte {
	if ch == '}' {
		return '{'
	}
	if ch == ']' {
		return '['
	}
	return '('
}

func isLeft(ch byte) bool {
	return ch == '{' || ch == '[' || ch == '('
}
