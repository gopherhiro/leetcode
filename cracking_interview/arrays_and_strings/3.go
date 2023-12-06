package arrays_and_strings

import "strings"

// 题目：URL 化
// 思路：构造字符串
// 复杂度：time O(n), space O(1)
func replaceSpaces(s string, length int) string {
	var builder strings.Builder
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			builder.WriteString("%20")
		} else {
			builder.WriteByte(s[i])
		}
	}
	return builder.String()
}
