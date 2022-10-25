package string

import (
	"strconv"
	"strings"
)

// 38. Count and Say
// 38. 外观数列
// 思路：字符串遍历&构造
// time O(n*m) space O(m)
func countAndSay(n int) string {
	if n <= 0 {
		return "-1"
	}
	res := "1"
	for i := 1; i < n; i++ {
		res = build(res)
	}
	return res
}

func build(s string) string {
	builder := &strings.Builder{}
	p := 0
	for p < len(s) {
		// get char and count
		ch := s[p]
		count := 0
		for p < len(s) && s[p] == ch {
			p++
			count++
		}
		// a pair of number count
		pair := strconv.Itoa(count) + string(ch)
		builder.WriteString(pair)
	}
	return builder.String()
}
