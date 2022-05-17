package slidingwindow

// 187. Repeated DNA Sequences
// 187. 重复的DNA序列
// 思路：sliding window
// time O(N) space O(N)
func findRepeatedDnaSequencesH(s string) (ans []string) {
	if len(s) == 0 {
		return
	}
	const subLen = 10
	submap := make(map[string]int, 0)
	left, right := 0, 0
	for right < len(s) {
		// 右扩窗口
		right++
		for right-left == subLen {
			// 更新结果
			sub := s[left:right]
			// 出现不止一次
			if submap[sub] == 1 {
				ans = append(ans, sub)
			}
			submap[sub]++
			// 左缩窗口
			left++
		}
	}
	return
}

// 187. Repeated DNA Sequences
// 187. 重复的DNA序列
// 思路：暴力枚举: hashtable
// time O(N) space O(N)
func findRepeatedDnaSequences(s string) (ans []string) {
	if len(s) == 0 {
		return
	}
	const subLen = 10
	submap := make(map[string]int, 0)
	for i := 0; i <= len(s)-subLen; i++ {
		sub := s[i : i+subLen]
		// 出现不止一次
		if submap[sub] == 1 {
			ans = append(ans, sub)
		}
		submap[sub]++
	}
	return
}
