package string

import "strings"

// 1678. Goal Parser Interpretation
// 1678. 设计 Goal 解析器
// 思路：一次遍历
// time O(n) space O(1)
func interpret(cmd string) string {
	if len(cmd) == 0 {
		return ""
	}
	ans := &strings.Builder{}
	i := 0
	for i < len(cmd) {
		if cmd[i] == 'G' {
			ans.WriteString("G")
			i++
			continue
		}
		if cmd[i] == '(' && cmd[i+1] == ')' {
			ans.WriteString("o")
			i += 2
			continue
		}
		if cmd[i] == '(' && cmd[i+1] == 'a' && cmd[i+2] == 'l' && cmd[i+3] == ')' {
			ans.WriteString("al")
			i += 4
		}
	}
	return ans.String()
}
