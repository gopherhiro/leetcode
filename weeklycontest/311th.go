package main

// 6180. 最小偶倍数
// time O(n) space O(1)
func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	i, two := 1, 2
	for {
		multi := i * two
		if multi%n == 0 {
			break
		}
		i++
	}
	return i * two
}

// 6181. 最长的字母序连续子字符串的长度
// time O(n) space O(1)
func longestContinuousSubstring(s string) int {
	ans, count := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1]+1 {
			count++
		} else {
			ans = max(ans, count)
			count = 1
		}
	}
	ans = max(ans, count)
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
