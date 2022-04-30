package main

// 6051. 统计是给定字符串前缀的字符串数目
func countPrefixes(words []string, s string) int {
	count := 0
	for _, word := range words {
		if isPrefix(word, s) {
			count++
		}
	}
	return count
}

func isPrefix(substr, s string) bool {
	if len(substr) > len(s) {
		return false
	}

	i, j := 0, 0
	for i < len(substr) && j < len(s) {
		if substr[i] == s[j] {
			i++
			j++
		} else {
			return false
		}
	}

	if i == len(substr) {
		return true
	}
	return false
}
