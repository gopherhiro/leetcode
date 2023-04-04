package number

import "math"

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

	// 不在[-2^31..2^31-1] 范围内，返回0
	if sum > math.MaxInt32 || sum < math.MinInt32 {
		return 0
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
	// also ok, for x != 0
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
