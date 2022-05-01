package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func main() {
	a := "hello"
	b := "ll"
	n := strStr(a, b)
	fmt.Println(n)
}

// 28. Implement strStr()
// 28. 实现 strStr()
func strStr(haystack string, needle string) int {
	return 0
}

// 67. Add Binary
// 67. 二进制求和
// 思路：模拟加法
// time (max(m,n)) space O(max(m,n))
func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	ans := make([]byte, 0)
	carry := 0 // 进位
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		va, vb := 0, 0
		if i >= 0 {
			va = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			vb = int(b[j] - '0')
			j--
		}
		sum := va + vb + carry // 求和
		carry = sum / 2        // 进位
		curr := sum % 2        // 当前值
		res := byte(curr + '0')
		ans = append([]byte{res}, ans...)
	}
	// 如果最后有进位，则需要再加一位
	if carry > 0 {
		last := byte(carry + '0')
		ans = append([]byte{last}, ans...)
	}
	return string(ans)
}

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：sort 两个字符串排序后相等
// time O(N) space O(M)
func isAnagramP(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	a, b := []byte(s), []byte(t)
	// 排序
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	// 比较
	return string(a) == string(b)
}

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：hash table
// time O(N) space O(M)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 统计
	countS := make(map[byte]int, 0)
	countT := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++
	}
	// 比较
	for key, val := range countS {
		if countT[key] != val {
			return false
		}
	}
	return true
}

// 242. Valid Anagram
// 242. 有效的字母异位词
// 思路：hash table
// time O(N) space O(M)
func isAnagramO(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	// 统计
	countS := make(map[byte]int, 0)
	countT := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		countS[s[i]]++
		countT[t[i]]++
	}
	// 比较
	return reflect.DeepEqual(countS, countT)
}

// 169. Majority Element
// 169. 多数元素
// 思路：哈希表统计
// time O(N) space O(N)
func majorityElement(nums []int) (ans int) {
	hash := make(map[int]int, 0)
	n := len(nums)

	majorityCount := n / 2
	for i := 0; i < n; i++ {
		hash[nums[i]]++
		if hash[nums[i]] > majorityCount {
			ans = nums[i]
		}
	}
	return
}

// 169. Majority Element
// 169. 多数元素
// 思路：排序即可。
// time O(N) space O(N or 1)
func majorityElementO(nums []int) (ans int) {
	sort.Ints(nums)
	ans = nums[len(nums)/2]
	return
}

// 169. Majority Element
// 169. 多数元素
// 思路：Boyer-Moore 投票算法
// 1、遇到相同的就 +1，不同的就 -1。
// 2、减到 0 就换个数，重新计数。
// 3、总能找到最多的那个数。
// time O(N) space O(1)
func majorityElementP(nums []int) int {
	candidate, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

// 12. Integer to Roman
// 12. 整数转罗马数字
// 思路：模拟法
// time O(N) space O(N)
func intToRoman(num int) string {
	type romanNode struct {
		val int
		str string
	}
	romanNodes := []romanNode{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// 遍历寻找到 num 包含的范围
	// 获取对应字符拼接即可
	sb := strings.Builder{}
	for _, node := range romanNodes {
		for num >= node.val {
			sb.WriteString(node.str)
			num -= node.val
		}
		if num == 0 {
			break
		}
	}
	return sb.String()
}

// 13. Roman to Integer
// 13. 罗马数字转整数
// 思路：根据罗马数字与整数规律即可。
// time O(N) space O(1)
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans, n := 0, len(s)
	for i := 0; i < n; i++ {
		curr := romanMap[s[i]]
		// 若数字比它右侧的数字小，则将该数字的符号取反。
		if i < n-1 && curr < romanMap[s[i+1]] {
			curr = -curr
		}
		ans += curr
	}
	return ans
}

// 58. Length of Last Word
// 58. 最后一个单词的长度
// 思路：反向遍历
// time O(N) space O(1)
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastWordLen, i := 0, len(s)-1
	// 跳过空格
	for i >= 0 && s[i] == ' ' {
		i--
	}
	// 开始统计最后一个单词的长度
	// 遇到空格，表示最后一个单词结束
	for i >= 0 && s[i] != ' ' {
		lastWordLen++
		i--
	}
	return lastWordLen
}

// Golang 本身的语言特性
func lengthOfLastWordN(s string) int {
	t := strings.Fields(s)
	if len(t) == 0 {
		return 0
	}
	last := t[len(t)-1]
	return len(last)
}

// 14. Longest Common Prefix
// 14. 最长公共前缀
// 思路：先求两个字符串的公共前缀，然后遍历即可。
// 备注：横向遍历
// time O(s = m * n) space O(1)
func longestCommonPrefix(str []string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str[0]
	}

	// 默认第一个单词就是最长公共前缀
	prefix := str[0]

	// 逐个比较所有单词，获取最长公共前缀
	for i := 1; i < len(str); i++ {
		prefix = LCP(prefix, str[i])
		if len(prefix) == 0 {
			return ""
		}
	}
	return prefix
}

// 求两个字符串的公共前缀
// time O(N) space O(1)
func LCP(s1, s2 string) string {
	length := min(len(s1), len(s2))
	index := 0
	for index < length && s1[index] == s2[index] {
		index++
	}
	return s1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
