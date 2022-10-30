package BFS

// 784. Letter Case Permutation
// 784. 字母大小写全排列
// 思路：BFS
// time O(n*2^n) space O(n*2^n)
func letterCasePermutation(s string) []string {
	queue := []string{s}
	for i := 0; i < len(s); i++ {
		// 对 s[i] 的字符进行「大小写变换」，即得到一个结果，
		// 得到的结果直接存回队列即可。
		// 字母大小写的 ASCII 差为 32，
		// 异或 32 可进行字母大小写转换。
		if isNumber(s[i]) {
			continue
		}
		for _, str := range queue {
			chars := []byte(str)
			chars[i] = chars[i] ^ 32
			queue = append(queue, string(chars))
		}
	}
	return queue
}

func isNumber(char byte) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}
