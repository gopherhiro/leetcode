package math

import (
	"math"
)

func main() {

}

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

// 2119. A Number After a Double Reversal
// 2119. 反转两次的数字
// 思路：反转两次，对比即可。
// time O(N) space O(1)
func isSameAfterReversalsN(num int) bool {
	if num == 0 {
		return true
	}
	r1 := reverseN(num)
	r2 := reverseN(r1)
	return r2 == num
}

func isSameAfterReversals(num int) bool {
	// 如果一个数可以 mod 10
	// 反转两次后，必然不等于之前的数
	return num == 0 || num%10 != 0
}

// 7. Reverse Integer
// 7. 整数反转
// 思路：取余法
// time O(N) space O(1)
func reverseN(x int) int {
	if x == 0 {
		return 0
	}
	sum := 0
	for x != 0 {
		remain := x % 10
		sum = sum*10 + remain
		x = x / 10
	}
	return sum
}

// 7. Reverse Integer
// 7. 整数反转
// 思路：取余法
// time O(N) space O(1)
func reverse(x int) int {
	if x == 0 {
		return 0
	}

	// 负数标记
	isNeg := false
	if x < 0 {
		isNeg = true
		x = -x
	}

	sum := 0
	for x > 0 {
		remain := x % 10
		sum = sum*10 + remain
		x = x / 10
	}
	// 如果是负数，则还原负号。
	if isNeg {
		sum = -sum
	}

	// 不在[-2^31..2^31-1] 范围内，返回0
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
	}

	return sum
}

// 292. Nim 游戏
func canWinNim(n int) bool {
	return n%4 != 0
}

// 319. 灯泡开关
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}
