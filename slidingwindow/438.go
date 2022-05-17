package slidingwindow

// 438. Find All Anagrams in a String
// 438. 找到字符串中所有字母异位词
// 思路：滑动窗口
// time O(N) space O(N)
func findAnagrams(s string, p string) (ans []int) {
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return
	}
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)

	for i := 0; i < len(p); i++ {
		e := p[i]
		need[e]++
	}

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// 右扩窗口
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++

		// 根据条件，左缩窗口
		for right-left == len(p) {
			// 更新结果集
			if valid == len(need) {
				ans = append(ans, left)
			}
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
	return
}
