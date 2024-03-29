package main

// 100245. 每个字符最多出现两次的最长子字符串
func maximumLengthSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	window := make(map[byte]int, 0)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++

		// 窗口内有重复的字符
		// 进行左缩，直到没有重复字符为止
		for window[c] > 2 {
			d := s[left]
			left++
			window[d]--
		}
		// 保证没有重复字符，再进行统计
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
