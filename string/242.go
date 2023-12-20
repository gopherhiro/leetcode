package string

import (
	"reflect"
	"sort"
)

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：sort 两个字符串排序后相等
// time O(N+M) space O(1)
func isAnagramN(s string, t string) bool {
	return countWord(s) == countWord(t)
}

func countWordN(s string) [26]rune {
	counter := [26]rune{}
	for _, v := range s {
		counter[v-'a']++
	}
	return counter
}

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：sort 两个字符串排序后相等
// time O(N) space O(M)
func isAnagramP(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a, b := []byte(s), []byte(t)
	// 排序
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	// 比较
	return string(a) == string(b)
}

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：hash table
// time O(N) space O(M)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 统计
	countS := make(map[byte]int, 0)
	countT := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++
	}
	// 比较
	for key, val := range countS {
		if countT[key] != val {
			return false
		}
	}
	return true
}

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：hash table
// time O(N) space O(M)
func isAnagramO(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 统计
	countS := make(map[byte]int, 0)
	countT := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++
	}
	// 比较
	return reflect.DeepEqual(countS, countT)
}
