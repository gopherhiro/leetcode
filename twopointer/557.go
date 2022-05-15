package twopointer

import "strings"

// 557. Reverse Words in a String III
// 557. 反转字符串中的单词 III
// 思路：以空格分割句子成单词数组，
// 遍历单词数组，reverse每一个单词。
// time O(N), space O(N)
func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	words := strings.Split(s, " ")
	newWords := make([]string, 0)
	for _, word := range words {
		t := []byte(word)
		reverseString(t)
		newWords = append(newWords, string(t))
	}
	return strings.Join(newWords, " ")
}
