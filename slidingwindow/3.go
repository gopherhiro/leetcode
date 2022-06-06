package slidingwindow

// 3. Longest Substring Without Repeating Characters
// 3. 无重复字符的最长子串
// 思路：sliding window
// time O(N) space O(N)
func lengthOfLongestSubstring(s string) int {
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
		for window[c] > 1 {
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

// 字符串 s 的所有无重复子串
// 思路：sliding window
// time O(N) space O(N)
func noDubSubstr(s string) (ans []string) {
	if len(s) == 0 {
		return
	}
	window := make(map[byte]int, 0)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++

		// 窗口内有重复的字符
		// 进行左缩，直到没有重复字符为止
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		ans = append(ans, s[left:right])
	}
	return
}
