package string

import "strconv"

// 415. 字符串相加
// 思路：模拟加法
// time O(max(m,n)) space O(1)
func addStrings(num1 string, num2 string) string {
	if len(num1) == 0 {
		return num2
	}

	if len(num2) == 0 {
		return num1
	}

	var result string
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	for i >= 0 || j >= 0 || carry > 0 {
		var vi, vj int

		if i >= 0 {
			vi = int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			vj = int(num2[j] - '0')
			j--
		}

		sum := vi + vj + carry // 求和
		carry = sum / 10       // 进位
		curr := sum % 10       // 当前值

		result = strconv.Itoa(curr) + result
	}
	return result
}
