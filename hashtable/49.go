package hashtable

// 49. Group Anagrams
// 49. 字母异位词分组
// 思路：哈希表
// time O(nk) space O(nk)
func groupAnagrams(words []string) [][]string {
	// 字符出现频次一致的单词，分成一个组
	m := make(map[[26]byte][]string, 0)
	for _, word := range words {
		key := countWord(word)
		m[key] = append(m[key], word)
	}
	// 把分组结果 append 在一起即可。
	answer := make([][]string, 0, len(m))
	for _, group := range m {
		answer = append(answer, group)
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
