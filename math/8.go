package math

import "math"

// 8. String to Integer (atoi)
// 8. 字符串转换整数 (atoi)
// 思路：字符串处理 + 数字构建
// time O(N) space O(1)
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	sum, i, n := 0, 0, len(s)

	// 跳过前导空格
	for i < n && s[i] == ' ' {
		i++
	}

	// 判断正负
	// 使用sign 标记正负数（很不错的技巧）
	sign := 1 // 默认为正数
	if i < n {
		if s[i] == '-' {
			sign = -1
			i++
		} else if s[i] == '+' {
			sign = 1
			i++
		}
	}

	// 数字部分，累加求和
	for i < n && isNumeric(s[i]) {
		number := int(s[i] - '0')
		sum = sum*10 + number
		// 累加过程中，超范围，直接返回
		if sign*sum > math.MaxInt32 {
			return math.MaxInt32
		}
		if sign*sum < math.MinInt32 {
			return math.MinInt32
		}
		i++
	}

	return sign * sum
}

// 判断是否是字符或数字
func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
