package parentheses

// 921. Minimum Add to Make Parentheses Valid
// 921. 使括号有效的最少添加
// 思路：Balance
// time O(N), space O(1)
func toMakeValid(s string) int {
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
