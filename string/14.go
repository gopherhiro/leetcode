package string

// 14. Longest Common Prefix
// 14. 最长公共前缀
// 思路：先求两个字符串的公共前缀，然后遍历即可。
// 备注：横向遍历
// time O(s = m * n) space O(1)
func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}

	// 默认第一个单词就是最长公共前缀
	prefix := str[0]

	// 逐个比较所有单词，获取最长公共前缀
	for i := 1; i < len(str); i++ {
		prefix = LCP(prefix, str[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

// 求两个字符串的公共前缀
// time O(N) space O(1)
func LCP(s1, s2 string) string {
	length := min(len(s1), len(s2))
	index := 0
	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
