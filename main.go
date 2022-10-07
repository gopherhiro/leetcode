package main

import (
	"fmt"
)

func main() {
	s := "aabb"
	r := firstUniqChar2(s)
	fmt.Println(r)
}

// 387. First Unique Character in a String
// 387. 字符串中的第一个唯一字符
// 思路：哈希表存储索引
// time O(n) time O(1)
func firstUniqChar2(s string) int {
	ht := make(map[rune]int, 0)
	for i, ch := range s {
		if _, ok := ht[ch]; ok {
			ht[ch] = -1
		} else {
			ht[ch] = i
		}
	}
	first := len(s)
	for _, index := range ht {
		if index == -1 {
			continue
		}
		first = min(first, index)
	}
	if first == len(s) {
		first = -1
	}
	return first
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 387. First Unique Character in a String
// 387. 字符串中的第一个唯一字符
// 思路：哈希表
// time O(n) time O(1)
func firstUniqChar(s string) int {
	ht := make(map[rune]int, 0)
	for _, ch := range s {
		ht[ch]++
	}
	for i, ch := range s {
		if ht[ch] == 1 {
			return i
		}
	}
	return -1
}
