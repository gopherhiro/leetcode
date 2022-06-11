package main

import "fmt"

func main() {
	s := "[] {} ()"
	r := parenthesesIsValid(s)
	fmt.Println(r)
}

// 20. Valid Parentheses
// 20. 有效的括号
// time O(N), space O(N)
func parenthesesIsValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if v == ' ' {
			continue
		}
		if v == '(' || v == '[' || v == '{' {
			// 入栈
			stack = append(stack, v)
			continue
		}
		// 取栈顶元素
		// 如果 s 不为空，但 stack 为空时，不合法。
		if len(stack) == 0 {
			return false
		}
		top := stack[len(stack)-1]
		valid := top == '(' && v == ')' || top == '{' && v == '}' || top == '[' && v == ']'
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
