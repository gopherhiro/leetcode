package backtracking

import "strings"

// 22. Generate Parentheses
// 22. 括号生成
func generateParenthesis(n int) (res []string) {
	if n == 0 {
		return
	}

	track := make([]string, 0)

	// left 记录还可以使用多少个左括号，
	// right 记录还可以使用多少个右括号。
	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		// 剪枝：不合法的结果
		// 合法的括号组合中：剩余的左括号一定是 <= 右括号的
		if left > right {
			return
		}
		if left < 0 || right < 0 {
			return
		}
		// 结束条件
		// 当左右括号刚好用完，得到一个合法的括号组合。
		if left == 0 && right == 0 {
			tmp := make([]string, len(track))
			copy(tmp, track)
			res = append(res, strings.Join(tmp, ""))
		}

		// 选择列表，进行路径选择
		// 尝试选择左括号
		track = append(track, "(")
		backtrack(left-1, right)
		track = track[:len(track)-1]

		// 尝试选择右括号
		track = append(track, ")")
		backtrack(left, right-1)
		track = track[:len(track)-1]
	}
	// 刚开始时，左括号和右括号相等
	backtrack(n, n)
	return
}
