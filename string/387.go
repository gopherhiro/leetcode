package string

// 387. First Unique Character in a String
// 387. 字符串中的第一个唯一字符
// 思路：哈希表统计频次
// time O(n) time O(1)
func firstUniqChar(s string) int {
	ht := make(map[rune]int, 0)
	// 第一次遍历：统计频次
	for _, ch := range s {
		ht[ch]++
	}
	// 第二次遍历：找寻第一个频次为 1 的字符下标即可
	for i, ch := range s {
		if ht[ch] == 1 {
			return i
		}
	}
	return -1
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
