package string

// 2273. Find Resultant Array After Removing Anagrams
// 2273. 移除字母异位词后的结果数组
// 思路：模拟
// time O(nk) space O(n)
func removeAnagrams(words []string) []string {
	// 把 words[0] 预存到 answer 中
	answer := []string{words[0]}
	for i := 1; i < len(words); i++ {
		// 依次比较，如果相邻元素是「字母异位词」，跳过。
		// 相邻元素非「字母异位词」，则添加进入结果集中。
		if isAnagrams(words[i], words[i-1]) {
			continue
		}
		answer = append(answer, words[i])
	}
	return answer
}

func isAnagrams(a, b string) bool {
	return countWord(a) == countWord(b)
}

// 2273. Find Resultant Array After Removing Anagrams
// 2273. 移除字母异位词后的结果数组
// 思路：哈希表
// time O(nk) space O(n)
func removeAnagrams2(words []string) []string {
	answer := make([]string, 0)
	seen := make(map[[26]byte]bool, 0)
	for _, word := range words {
		key := countWord(word)
		if seen[key] {
			continue
		}
		answer = append(answer, word)
		seen[key] = true
	}
	return answer
}

// 编码单词频次作为哈希表的 key
func countWord(s string) [26]byte {
	count := [26]byte{}
	for _, v := range s {
		count[v-'a']++
	}
	return count
}
