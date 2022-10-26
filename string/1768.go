package string

import "strings"

// 1768. Merge Strings Alternately
// 1768. 交替合并字符串
// 思路：two pointer
// time O(n) space O(1)
func mergeAlternately(word1 string, word2 string) string {
	answer := &strings.Builder{}
	// i,j 分别指向两个单词的初始位置
	// 依次遍历 word1 and word2 中的字符 & 拼接到 answer 中
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		answer.WriteString(string(word1[i]))
		answer.WriteString(string(word2[j]))
		i++
		j++
	}
	// 较长的单词的末尾字符，统一拼接到 answer 中
	if i < len(word1) {
		answer.WriteString(word1[i:])
	}
	if j < len(word2) {
		answer.WriteString(word2[j:])
	}
	// 返回 answer 即可
	return answer.String()
}
