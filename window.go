package main

import "fmt"

func main() {
	s := "abcabcbb"
	//t := "ab"
	str := lengthOfLongestSubstring(s)
	fmt.Println(str)
}

func lengthOfLongestSubstring(s string) int {
	var res int
	window := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		right++
		window[c]++

		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}

		if right-left > res {
			res = right-left
		}

	}
	return res
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findAnagrams(s string, p string) []int {
	var res []int
	if len(s) == 0 || len(p) == 0 {
		return res
	}
	if len(p) > len(s) {
		return res
	}

	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}

	left, right, valid := 0, 0, 0
	for right < len(s) {

		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for right-left >= len(p) {

			if valid == len(need) {
				res = append(res, left)
			}

			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}

	}
	return res
}

func checkInclusion(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(t) > len(s) {
		return false
	}

	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// 右扩窗口
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		// left缩时机是窗口大小等于len(t)
		for right-left >= len(t) {
			if valid == len(need) {
				return true
			}

			// 左缩窗口
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}

		}

	}
	return false
}

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	if len(t) > len(s) {
		return ""
	}

	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	// 记录最小子串的起始索引及长度
	start, length := 0, len(s)+1

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// c 是即将移入窗口的字符
		c := s[right]
		// 右扩窗口
		right++
		// 窗口内数据统计&更新
		if _, ok := need[c]; ok {
			window[c]++
			if need[c] == window[c] {
				valid++
			}
		}

		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
			}
			// d是即将移除窗口的字符
			d := s[left]
			// 左缩窗口
			left++
			// 窗口内数据统计&更新
			if _, ok := need[d]; ok {
				if need[d] == window[d] {
					valid--
				}
				window[d]--
			}
		}

	}

	// 返回最小覆盖子串
	if length == len(s)+1 {
		return ""
	}
	return s[start : start+length]
}
