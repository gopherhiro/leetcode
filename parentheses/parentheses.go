package parentheses

// 1541. Minimum Insertions to Balance a Parentheses String
// 1541. 平衡括号字符串的最少插入次数
// time O(N), space O(1)
func minInsertions(s string) int {
	if len(s) == 0 {
		return 0
	}

	res, rightNeed := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 一个左括号对应两个右括号
			rightNeed += 2

			if rightNeed%2 == 1 {
				// 插入一个右括号
				res++
				// 对右括号的需求减一
				rightNeed--
			}
		}

		if s[i] == ')' {
			rightNeed--
			// 右括号太多了
			if rightNeed == -1 {
				// 需要插入一个左括号
				res++
				// 同时，对右括号的需求变为 1
				rightNeed = 1
			}
		}

	}
	return res + rightNeed
}

// 921. Minimum Add to Make Parentheses Valid
// 921. 使括号有效的最少添加
// 思路：Balance
// time O(N), space O(1)
func ToMakeValid(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans, bal := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			bal += 1
		} else {
			bal -= 1
		}
		// It is guaranteed bal >= -1
		if bal == -1 {
			ans++
			bal++
		}

	}

	return ans + bal
}

// 921. Minimum Add to Make Parentheses Valid
// 921. 使括号有效的最少添加
// 思路：统计左右括号的需求量，加起来即可。
// time O(N), space O(1)
func minAddToMakeValid(s string) int {
	if len(s) == 0 {
		return 0
	}
	// 统计左右括号的需求量
	leftNeed, rightNeed := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			// 右括号的需求量：+1
			rightNeed++
		}
		if s[i] == ')' {
			// 右括号的需求量：-1
			rightNeed--

			// 右括号太多了
			if rightNeed == -1 {
				rightNeed = 0

				// 需要左括号
				leftNeed++
			}
		}

	}

	return leftNeed + rightNeed
}

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
