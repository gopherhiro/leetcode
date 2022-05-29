package string

import (
	"math"
)

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

// 2287. Rearrange Characters to Make Target String
// 6078. 重排字符形成目标字符串
// 策略：哈希统计+木桶效应
// time O(N) space O(N)
func rearrangeCharactersN(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	// 统计 s 中的字符
	sCount := make(map[rune]int, 0)
	for _, v := range s {
		sCount[v]++
	}
	// 统计 t 中的字符
	tCount := make(map[rune]int, 0)
	for _, v := range t {
		tCount[v]++
	}
	ans := math.MaxInt16
	for ch := range tCount {
		// s 中只要不包含 t 中的任何一个字符，则不能形成 t 的副本。
		// 直接 return	0
		if sCount[ch] == 0 {
			return 0
		}
		// 木桶效应：桶能装多少水取决于它最短的那块木板。
		ans = min(ans, sCount[ch]/tCount[ch])
	}
	if ans == math.MaxInt16 {
		return 0
	}
	return ans
}
