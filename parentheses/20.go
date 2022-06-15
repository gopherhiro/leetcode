package parentheses

// 20. Valid Parentheses
// 20. 有效的括号
// time O(N), space O(N)
func parenthesesIsValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			// 入栈
			stack = append(stack, s[i])
			continue
		}
		// 取栈顶元素
		// 如果 s 不为空，但 stack 为空时，不合法。
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		valid := top == '(' && s[i] == ')' || top == '{' && s[i] == '}' || top == '[' && s[i] == ']'
		if valid {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	// 遍历完 s，栈空，合法。
	if len(stack) == 0 {
		return true
	}
	return false
}

// 20. Valid Parentheses
// 20. 有效的括号
// time O(N), space O(N)
func pisValid(s string) bool {
	if len(s) == 0 {
		return false
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) != 0 && leftOf(s[i]) == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			// 和最近的左括号不匹配
			return false
		}
	}

	// 遍历完s，所有的左括号都被匹配了。
	if len(stack) == 0 {
		return true
	}
	return false
}

func leftOf(ch byte) byte {
	if ch == ')' {
		return '('
	}
	if ch == '}' {
		return '{'
	}
	return '['
}

// 20. Valid Parentheses
// 20. 有效的括号
// time O(N), space O(N)
func isValid06(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if isLeft(s[i]) {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		if leftOf(s[i]) == top {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	// 如果栈空，则是有效括号
	if len(stack) == 0 {
		return true
	}
	return false
}

func leftOf06(ch byte) byte {
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
