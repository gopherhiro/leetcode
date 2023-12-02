package arrays_and_strings

import "sort"

// 题目：判定是否互为字符重排
// 思路：检查两个字符串的字符数是否相同
// 复杂度：time O(m+n), space O(c)
// 使用哈希表将其每个字符映射到其字符出现的次数。
// 遍历第一个字符串，增加字符出现个数。
// 遍历第二个字符串，减少字符出现个数。
// 如果两者互为重排，则哈希表中统计的字符个数最终将为 0。
// 若值为负值就提早终止，因为：
// 1. 字符串长度相同，增加的次数与减少的次数也相同。
// 2. 若哈希表无负值， 则表示两字符串互为重排。

func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	letters := make(map[rune]int)
	for _, ch := range s1 {
		letters[ch]++
	}
	for _, ch := range s2 {
		letters[ch]--
		if letters[ch] < 0 {
			return false
		}
	}
	return true
}

// 题目：判定是否互为字符重排
// 思路：排序
// 复杂度：time O(nlogn), space O(logn)

func CheckPermutationS(s1 string, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)

	sort.Slice(b1, func(i, j int) bool {
		return b1[i] < b1[j]
	})

	sort.Slice(b2, func(i, j int) bool {
		return b2[i] < b2[j]
	})

	return string(b1) == string(b2)
}
