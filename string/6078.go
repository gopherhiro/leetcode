package string

// 2287. Rearrange Characters to Make Target String
// 6078. 重排字符形成目标字符串
// 策略：桶排序之木桶效应
// time O(N) space O(N)
func rearrangeCharacters(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	sCount := counter(s)
	tCount := counter(t)
	// 字符串 s 中最多包含 len(s) 个 字符串 t。
	ans := len(s)
	for ch, count := range sCount {
		// t 的统计桶中无字符，跳过。
		if tCount[ch] == 0 {
			continue
		}
		// 木桶效应：桶能装多少水取决于它最短的那块木板。
		ans = min(ans, count/tCount[ch])
	}
	return ans
}

func counter(s string) []int {
	buckets := make([]int, 26)
	if len(s) == 0 {
		return buckets
	}
	for _, v := range s {
		buckets[v-'a']++
	}
	return buckets
}
