package slidingwindow

import "math"

// 76. Minimum Window Substring
// 76. 最小覆盖子串
// 思路：sliding window (two pointer)
// time O(N) space O(N)
func minWindow(s string, t string) string {
	if len(s) == 0 || len(s) < len(t) {
		return ""
	}
	// need 记录字符串 t 中的字符出现的次数
	// window 记录窗口中对应 t 中的字符出现的次数
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)

	// 构建所需字符数的统计
	for i := 0; i < len(t); i++ {
		e := t[i]
		need[e]++
	}

	// 记录最小覆盖子串：起始索引、长度
	start, minLen := 0, math.MaxInt32

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// 加入元素，更新统计，右扩窗口
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++

		// need 中的所有字符统计数都满足时
		// valid == len(need)
		for valid == len(need) {

			// 更新[最小覆盖子串]结果
			subLen := right - left
			if subLen < minLen {
				// 更新最小覆盖子串：起始索引，长度
				start = left
				minLen = subLen
			}

			// 左缩窗口：移出元素，更新字符数据统计
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}
