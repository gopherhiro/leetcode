package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "AbCdEfGhIjK"
	r := greatestLetter(s)
	fmt.Println(r)
}

// 5242. 兼具大小写的最好英文字母
// 策略：哈希表 + 遍历
// time O(nlogn) time O(n)
func greatestLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	ht := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		ht[s[i]]++
	}

	ans := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if isLower(s[i]) {
			tmp := toUpper(s[i])
			if _, ok := ht[tmp]; ok {
				ans = append(ans, tmp)
			}
			continue
		}
		if isUpper(s[i]) {
			tmp := toLower(s[i])
			if _, ok := ht[tmp]; ok {
				ans = append(ans, s[i])
			}
			continue
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		return ans[i] > ans[j]
	})
	if len(ans) == 0 {
		return ""
	}
	return string(ans[0])
}

func toLower(ch byte) byte {
	return ch + 32
}

func toUpper(ch byte) byte {
	return ch - 32
}

func isLower(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}

func isUpper(ch byte) bool {
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	return false
}
