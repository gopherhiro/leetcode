package string

// 2351. First Letter to Appear Twice
// 2351. 第一个出现两次的字母
// 思路：哈希表
// time O(n) space O(n)
func repeatedCharacter(s string) byte {
	seen := make(map[rune]bool, 0)
	for _, ch := range s {
		if seen[ch] {
			return byte(ch)
		}
		seen[ch] = true
	}
	return byte(' ')
}

// 2351. First Letter to Appear Twice
// 2351. 第一个出现两次的字母
// 思路：哈希表
// time O(n) space O(n)
func repeatedCharacter2(s string) byte {
	seen := make(map[rune]int, 0)
	for _, ch := range s {
		seen[ch]++
		if seen[ch] == 2 {
			return byte(ch)
		}
	}
	return byte(' ')
}
