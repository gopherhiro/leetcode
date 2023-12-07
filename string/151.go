package string

import "strings"

// 151. 反转字符串中的单词
// 思路：字符串分割&反转&合并
// time O(N) space O(N)
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	newWords := make([]string, 0)
	words := strings.Split(s, " ")
	for _, word := range words {
		if word == "" {
			continue
		}
		newWords = append(newWords, word)
	}
	reverse(newWords)
	return strings.Join(newWords, " ")
}

func reverse(s []string) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
