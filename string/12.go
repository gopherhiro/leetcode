package string

import "strings"

// 12. Integer to Roman
// 12. 整数转罗马数字
// 思路：模拟法
// time O(N) space O(N)
func intToRoman(num int) string {
	type romanNode struct {
		val int
		str string
	}
	romanNodes := []romanNode{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// 遍历寻找到 num 包含的范围
	// 获取对应字符拼接即可
	sb := strings.Builder{}
	for _, node := range romanNodes {
		for num >= node.val {
			sb.WriteString(node.str)
			num -= node.val
		}
		if num == 0 {
			break
		}
	}
	return sb.String()
}
