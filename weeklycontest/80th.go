package main

import (
	"fmt"
)

func main() {
	s := []int{5, 2, 6, 8, 9, 7}
	k := int64(50)
	r := countSubarrays(s, k)
	fmt.Println(r)
}

// 6098. 统计得分小于 K 的子数组数目
// 思路：滑动窗口
// 备注：过不了全部case
func countSubarrays(nums []int, k int64) int64 {
	windows := make([]int, 0)
	right := 0
	ans := 0
	for right < len(nums) {
		n := nums[right]
		windows = append(windows, n)
		if n < int(k) {
			ans++
		}
		right++
		for scores(windows) >= int(k) {
			windows = windows[1:]

		}
		if scores(windows) < int(k) && len(windows) > 1 {
			ans++
		}

	}
	return int64(ans)
}

func scores(w []int) int {
	sum := 0
	for _, v := range w {
		sum += v
	}
	return sum * len(w)
}

// 6096. 咒语和药水的成功对数
// 策略：暴力模拟
// time O(N^2)，超时
func successfulPairs(s []int, p []int, succ int64) []int {
	ans := make([]int, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			if s[i]*p[j] >= int(succ) {
				ans[i]++
			}
		}
	}
	return ans
}

// 6095. 强密码检验器 II
// 策略：利用 ASCII 码
// time O(N) space O(1)
func strongPasswordCheckerII(s string) bool {
	if len(s) < 8 {
		return false
	}
	var lower, upper, number, spc, existcon bool
	for i := 0; i < len(s); i++ {
		// lower
		if s[i] >= 97 && s[i] <= 122 {
			lower = true
		}
		// upper
		if s[i] >= 65 && s[i] <= 90 {
			upper = true
		}
		// number
		if s[i] >= 48 && s[i] <= 57 {
			number = true
		}
		// special character
		if hasSpecialChar(s[i]) {
			spc = true
		}
		if i+1 < len(s) && s[i] == s[i+1] {
			existcon = true
		}
	}

	if lower && upper && number && spc && !existcon {
		return true
	}
	return false
}

func hasSpecialChar(ch byte) bool {
	hashtable := map[byte]struct{}{
		'!': struct{}{},
		'@': struct{}{},
		'#': struct{}{},
		'$': struct{}{},
		'%': struct{}{},
		'^': struct{}{},
		'&': struct{}{},
		'*': struct{}{},
		'(': struct{}{},
		')': struct{}{},
		'-': struct{}{},
		'+': struct{}{},
	}
	if _, ok := hashtable[ch]; ok {
		return true
	}
	return false
}
