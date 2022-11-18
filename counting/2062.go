package counting

import "strings"

// 2062. Count Vowel Substrings of a String
// 2062. 统计字符串中的元音子字符串
// 思路：位运算
// time O(n^2) space O(1)
// 11111 ==> aeiou
func countVowelSubstrings(w string) int {
	n := len(w)
	if n == 0 {
		return 0
	}

	s := "aeiou"

	count := 0
	for i := 0; i < n; i++ {
		tmp := 0
		for j := i; j < n; j++ {
			if !strings.ContainsRune(s, rune(w[j])) {
				break
			}
			if w[j] == 'a' {
				tmp |= 1
			}
			if w[j] == 'e' {
				tmp |= 1 << 1
			}
			if w[j] == 'i' {
				tmp |= 1 << 2
			}
			if w[j] == 'o' {
				tmp |= 1 << 3
			}
			if w[j] == 'u' {
				tmp |= 1 << 4
			}
			if tmp == 1<<5-1 {
				count++
			}
		}
	}
	return count
}

// 2062. Count Vowel Substrings of a String
// 2062. 统计字符串中的元音子字符串
// 思路：brute force
// time O(n^2) space O(n^2)
func countVowelSubstringsB(w string) int {
	n := len(w)
	if n == 0 {
		return 0
	}

	// hash table
	valid := make(map[byte]bool, 0)
	vows := []byte{'a', 'e', 'i', 'o', 'u'}
	for _, v := range vows {
		valid[v] = true
	}

	count := 0
	for i := 0; i < n; i++ {
		// the set of no repeat element
		set := make(map[byte]bool, 0)
		for j := i; j < n; j++ {
			if !valid[w[j]] {
				break
			}
			set[w[j]] = true
			if len(set) == 5 {
				count++
			}
		}
	}
	return count
}
