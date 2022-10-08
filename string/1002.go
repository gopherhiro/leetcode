package string

import "math"

// 1002. Find Common Characters
// 1002. 查找共用字符
// 思路：counter
// time O(mn) space O(1)
func commonChars(words []string) []string {
	// 初始化最小频次统计数组
	minfreq := make([]int, 26)
	for i := 0; i < 26; i++ {
		minfreq[i] = math.MaxInt16
	}

	for _, word := range words {
		// 遍历字符数组，统计字符频次
		freq := make([]int, 26)
		for _, v := range word {
			freq[v-'a']++
		}
		// 维护单个字符，在所有字符串中出现的最小频次
		for i, count := range minfreq {
			minfreq[i] = min(freq[i], count)
		}
	}
	// 根据字符最小频次，构造共用字符
	ans := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < minfreq[i]; j++ {
			ans = append(ans, string(rune(i+'a')))
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
