package main

import "fmt"

func main() {
	s := "zabc"
	r := longestContinuousSubstring(s)
	fmt.Println(r)
}

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
