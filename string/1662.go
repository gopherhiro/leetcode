package string

// 1662. Check If Two String Arrays are Equivalent
// 1662. 检查两个字符串数组是否相等
// 思路：two pointer | 把两个字符数组，看作两个字符串。
// time O(n+m) space O(1)
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	i, j := 0, 0 // pointing to the current string in the array words
	x, y := 0, 0 // pointing to the current characters in the strings
	for i < len(word1) && j < len(word2) {
		// 对应字符不相等，则直接返回 false
		if word1[i][x] != word2[j][y] {
			return false
		}
		// 对应字符相等，则继续后续比较
		// 如果到达当前单词末尾，则跳到下一个单词比较
		// 切换到下一个单词是新单词，所以其字符索引需要重置为 0
		x, y = x+1, y+1
		if x == len(word1[i]) {
			i++
			x = 0
		}
		if y == len(word2[j]) {
			j++
			y = 0
		}
	}
	// 都同时达到字符串末尾，则两者表示的字符串相等。
	return i == len(word1) && j == len(word2)
}

// 1662. Check If Two String Arrays are Equivalent
// 1662. 检查两个字符串数组是否相等
// 思路：模拟拼接，比较
// time O(nk) space O(nk)
func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	s1, s2 := "", ""
	for _, word := range word1 {
		s1 += word
	}
	for _, word := range word2 {
		s2 += word
	}
	return s1 == s2
}
