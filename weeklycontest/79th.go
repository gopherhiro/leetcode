package main

import (
	"fmt"
	"strings"
)

func main() {
	messages := []string{"How is leetcode for everyone", "Leetcode is useful for practice"}
	senders := []string{"Bob", "Charlie"}
	r := largestWordCount(messages, senders)
	fmt.Println(r)
}

// 6083. 判断一个数的数字计数是否等于数位的值
// 策略：哈希表+遍历
// time O(N) space O(N)
func digitCount(num string) bool {
	ht := make(map[int]int, 0)
	for _, v := range num {
		value := int(v) - '0'
		ht[value]++
	}
	for i, v := range num {
		val := int(v) - '0'
		if val != ht[i] {
			return false
		}
	}
	return true
}

// 6084. 最多单词数的发件人
// 策略：哈希表 + 遍历
// time O(N) space O(N)
func largestWordCount(messages []string, senders []string) string {
	ht := make(map[string]int, 0)
	for i, sender := range senders {
		ht[sender] += wordCount(messages[i])
	}
	maxSender := senders[0]
	maxCount := ht[senders[0]]
	for sender, count := range ht {
		if count > maxCount {
			maxCount = count
			maxSender = sender
		}
		if count == maxCount {
			res := strings.Compare(sender, maxSender)
			if res > 0 {
				maxCount = count
				maxSender = sender
			}
		}

	}
	return maxSender
}

func wordCount(s string) int {
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}
	return count + 1
}
