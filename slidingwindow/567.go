package slidingwindow

// 567. Permutation in String
// 567. 字符串的排列
// 思路：sliding window (two pointer)
// time O(N) space O(N)
func checkInclusion(t string, s string) bool {
	if len(s) == 0 || len(s) < len(t) {
		return false
	}
	// need 记录字符串 t 中的字符出现的次数
	// window 记录窗口中对应 t 中的字符出现的次数
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)

	// 构建所需字符数的统计
	for i := 0; i < len(t); i++ {
		e := t[i]
		need[e]++
	}

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

		// t 是 s 的子串，右扩长度 == len(t)时，才停止
		// 可确定左侧窗口需要收缩
		for right-left == len(t) {
			// need 中的所有字符统计数都满足时
			// 找到的子串长度等于 t && need 的字符对齐
			// 则找到对应全排列子串
			if valid == len(need) {
				return true
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
	return false
}
