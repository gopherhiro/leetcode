package backtracking

// 784. Letter Case Permutation
// 784. 字母大小写全排列
// 思路：回溯
// time O(n*2^n) space O(n*2^n)
func letterCasePermutation(s string) []string {
	ans := make([]string, 0)

	var backtrack func(chars []byte, i int)
	backtrack = func(chars []byte, i int) {
		if i == len(s) {
			ans = append(ans, string(chars))
			return
		}

		// 数字，跳过大小写转换
		if isNumber(chars[i]) {
			backtrack(chars, i+1)
			return
		}

		chars[i] = chars[i] ^ 32
		backtrack(chars, i+1)

		chars[i] = chars[i] ^ 32
		backtrack(chars, i+1)
	}

	backtrack([]byte(s), 0)
	return ans
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}
