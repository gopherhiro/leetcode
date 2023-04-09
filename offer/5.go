package offer

import "strings"

// 剑指 Offer 05. 替换空格
// 思路：字符串操作
// time O(n) space O(n)
func replaceSpace(s string) string {
	if len(s) == 0 {
		return ""
	}
	var str strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			str.WriteString("%20")
		} else {
			str.WriteByte(s[i])
		}
	}
	return str.String()
}

func replaceSpace2(s string) string {
	if len(s) == 0 {
		return ""
	}
	str := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			str = append(str, '%')
			str = append(str, '2')
			str = append(str, '0')
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}
