package string

import "strings"

// 58. Length of Last Word
// 58. 最后一个单词的长度
// 思路：反向遍历
// time O(N) space O(1)
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastWordLen, i := 0, len(s)-1
	// 跳过空格
	for i >= 0 && s[i] == ' ' {
		i--
	}
	// 开始统计最后一个单词的长度
	// 遇到空格，表示最后一个单词结束
	for i >= 0 && s[i] != ' ' {
		lastWordLen++
		i--
	}
	return lastWordLen
}

// 58. Length of Last Word
// 58. 最后一个单词的长度
// 思路： 利用 Go 语言特性
func lengthOfLastWordN(s string) int {
	t := strings.Fields(s)
	if len(t) == 0 {
		return 0
	}
	last := t[len(t)-1]
	return len(last)
}
