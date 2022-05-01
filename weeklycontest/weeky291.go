package main

import (
	"leetcode/pkg"
	"math"
)

// 6050. 字符串的总引力
// 思路：sliding window
// time O(N^3) space O(N)
// 超时
func appealSum(s string) (ans int64) {
	if len(s) == 0 {
		return
	}
	left, right := 0, 0
	for right < len(s) {
		right++
		sub := s[left:right]
		ans += gravity(sub)
		if right == len(s) {
			left++
			right = left
		}
	}
	return
}

func gravity(s string) int64 {
	if len(s) == 0 {
		return 0
	}
	hash := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		hash[s[i]]++
	}
	return int64(len(hash))
}

// 6048. 必须拿起的最小连续卡牌数
// 思路：sliding window
// time O(N) space O(N)
func minimumCardPickup(cards []int) int {
	if len(cards) == 0 {
		return -1
	}
	ans := math.MaxInt32
	window := make(map[int]int, 0)
	left, right := 0, 0
	for right < len(cards) {
		c := cards[right]
		window[c]++
		right++

		// 窗口内有重复的数字
		// 进行左缩，直到没有重复数字为止
		for window[c] > 1 {
			ans = pkg.Min(ans, right-left)
			d := cards[left]
			left++
			window[d]--
		}
	}
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

// 6047. 移除指定数字得到的最大结果
// 策略：字符串操作
// time O(N^2) space O(1)
func removeDigit(number string, digit byte) string {
	max := ""
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			str := removeIndex(number, i)
			max = maxStr(max, str)
		}
	}
	return max
}

func maxStr(s1, s2 string) string {
	if len(s1) == 0 {
		return s2
	}
	if len(s2) == 0 {
		return s1
	}
	i := 0
	for i < len(s1) {
		if s1[i] > s2[i] {
			return s1
		}
		if s1[i] < s2[i] {
			return s2
		}
		// s1[i] == s2[i]
		i++
	}
	return s1
}

func removeIndex(number string, i int) string {
	if len(number) == 0 {
		return ""
	}
	if i < 0 || i >= len(number) {
		return number
	}

	prev := number[:i]
	last := number[i+1:]
	return prev + last
}
