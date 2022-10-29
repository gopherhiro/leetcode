package hashtable

import "unicode"

// 819. Most Common Word
// 819. 最常见的单词
// 思路：哈希表
// time O(m+n) space O(m+n)
func mostCommonWord(p string, banned []string) string {
	// 建立禁用词哈希表
	ban := make(map[string]bool, 0)
	for _, word := range banned {
		ban[word] = true
	}

	// 遍历段落，统计单词个数。
	freq := make(map[string]int, 0)
	maxFreq := 0
	var chars []byte
	for i := 0; i <= len(p); i++ {
		// 与其考虑过滤掉所有特殊字符
		// 不如关注我们所需要识别的字母
		// 循环条件 i <= len(p) 和 i < len(p) 是为了规避特殊情况：p = "Bob", banned = []
		if i < len(p) && unicode.IsLetter(rune(p[i])) {
			char := unicode.ToLower(rune(p[i]))
			chars = append(chars, byte(char))
			continue
		}
		if len(chars) == 0 {
			continue
		}
		// 得到一个单词
		// 不在禁用词列表中，统计其频次
		word := string(chars)

		if !ban[word] {
			freq[word]++
			maxFreq = max(maxFreq, freq[word])
		}
		// chars 重置为 nil，进行下一个单词的统计
		chars = nil
	}

	// 遍历频次统计表，查询频次最大的单词即可。
	for word, count := range freq {
		if count == maxFreq {
			return word
		}
	}
	return ""
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
