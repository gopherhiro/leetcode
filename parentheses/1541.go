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
