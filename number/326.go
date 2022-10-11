package math

import "math"

// 326. Power of Three
// 326. 3 的幂
// 思路：数学
// time O(logn) space O(1)
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}

// 326. Power of Three
// 326. 3 的幂
// 思路：打表
// time O(logn) space O(logn)
func isPowerOfThree2(n int) bool {
	if n <= 0 {
		return false
	}

	m := make(map[int]bool, 0)
	cur := 1
	m[cur] = true

	for cur <= math.MaxInt32/3 {
		cur *= 3
		m[cur] = true
	}

	if m[n] {
		return true
	}
	return false
}

// 326. Power of Three
// 326. 3 的幂
// 思路：打表2
// time O(1) space O(1)
func isPowerOfThree3(n int) bool {
	var m = map[int]bool{
		1:          true,
		3:          true,
		9:          true,
		27:         true,
		81:         true,
		243:        true,
		729:        true,
		2187:       true,
		6561:       true,
		19683:      true,
		59049:      true,
		177147:     true,
		531441:     true,
		1594323:    true,
		4782969:    true,
		14348907:   true,
		43046721:   true,
		129140163:  true,
		387420489:  true,
		1162261467: true,
	}
	if m[n] {
		return true
	}
	return false
}
