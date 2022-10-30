package DFS

// 784. Letter Case Permutation
// 784. 字母大小写全排列
// 思路：DFS
// time O(n*2^n) space O(n*2^n)
func letterCasePermutation(s string) []string {
	ans := make([]string, 0)

	var dfs func(chars []byte, i int)
	dfs = func(chars []byte, i int) {
		if i == len(s) {
			ans = append(ans, string(chars))
			return
		}

		// 数字，跳过转换
		if isNumber(chars[i]) {
			dfs(chars, i+1)
			return
		}

		// 如果是大写，则转成小写
		chars[i] = chars[i] ^ 32
		dfs(chars, i+1)

		// 如果是小写，则转成大写
		chars[i] = chars[i] ^ 32
		dfs(chars, i+1)
	}

	dfs([]byte(s), 0)
	return ans
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}
