package arrays_and_strings

// 题目：回文排列
// 思路：哈希表
// 根据字符串长度，「回文串」可分为两种情况：
// 1.「回文串」长度为偶数：所有不同字符的出现次数都为「偶数」；
// 2.「回文串」长度为奇数：位于中点的字符出现「奇数」次，其余字符出现「偶数」次；
// 因此，某字符串是回文串排列之一的「充要条件」为：
// 此字符串中，最多只有一种字符的出现次数为「奇数」，其余所有字符的出现次数都为「偶数」。
// 复杂度：time O(n), space O(n)
func canPermutePalindrome(s string) bool {
	letter := make(map[rune]int)
	for _, ch := range s {
		letter[ch]++
	}

	odd := 0
	for _, count := range letter {
		if count%2 == 1 {
			odd++
			if odd > 1 {
				return false
			}
		}
	}
	return true
}
