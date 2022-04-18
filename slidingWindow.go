package main

import (
	"math"
	"sort"
)

func main() {
}

// 187. Repeated DNA Sequences
// 187. 重复的DNA序列
// 思路：sliding window
// time O(N) space O(N)
func findRepeatedDnaSequencesH(s string) (ans []string) {
	if len(s) == 0 {
		return
	}
	const subLen = 10
	submap := make(map[string]int, 0)
	left, right := 0, 0
	for right < len(s) {
		// 右扩窗口
		right++
		for right-left == subLen {
			// 更新结果
			sub := s[left:right]
			// 出现不止一次
			if submap[sub] == 1 {
				ans = append(ans, sub)
			}
			submap[sub]++
			// 左缩窗口
			left++
		}
	}
	return
}

// 187. Repeated DNA Sequences
// 187. 重复的DNA序列
// 思路：暴力枚举: hashtable
// time O(N) space O(N)
func findRepeatedDnaSequences(s string) (ans []string) {
	if len(s) == 0 {
		return
	}
	const subLen = 10
	submap := make(map[string]int, 0)
	for i := 0; i <= len(s)-subLen; i++ {
		sub := s[i : i+subLen]
		// 出现不止一次
		if submap[sub] == 1 {
			ans = append(ans, sub)
		}
		submap[sub]++
	}
	return
}

// 220. Contains Duplicate III
// 220. 存在重复元素 III
// 思路：暴力枚举所有可能
// time O(N^2) space O(1)
func containsNearbyDup(nums []int, k, t int) bool {
	if len(nums) < 2 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i], nums[j]) <= t && abs(i, j) <= k {
				return true
			}
		}
	}
	return false
}

// 217. Contains Duplicate
// 217. 存在重复元素
// 思路：sort
// time O(N) space O(N)
func containsDuplicateN(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

// 217. Contains Duplicate
// 217. 存在重复元素
// 思路：hashtable
// time O(N) space O(N)
func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		elem := nums[i]
		hash[elem]++
		// 至少 2 次
		if hash[elem] >= 2 {
			return true
		}
	}
	return false
}

// 219. Contains Duplicate II
// 219. 存在重复元素 II
// 思路：sliding window
// time O(N) space O(N)
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	window := make(map[int]int, 0)
	left, right := 0, 0
	for right < len(nums) {
		c := nums[right]
		window[c]++
		right++

		// 存在重复值
		for window[c] > 1 {
			d := nums[left]
			left++
			window[d]--
			// 让 left 一直移至没有重复元素为止。
			if window[c] <= 1 {
				// 此时的坐标才是正确的
				// 计算坐标绝对值 & 与 k 比较
				if abs(left-1, right-1) <= k {
					return true
				}
			}
		}
	}

	return false
}

// 219. Contains Duplicate II
// 219. 存在重复元素 II
// 思路：hashtable
// time O(N) space O(N)
func containsNearbyDuplicateH(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	hash := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		elem := nums[i]
		if j, ok := hash[elem]; ok {
			if abs(i, j) <= k {
				return true
			}
		}
		hash[elem] = i
	}
	return false
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

// 3. Longest Substring Without Repeating Characters
// 3. 无重复字符的最长子串
// 思路：sliding window
// time O(N) space O(N)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	ans := 0
	window := make(map[byte]int, 0)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++

		// 窗口内有重复的字符
		// 进行左缩，直到没有重复字符为止
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		ans = max(ans, right-left)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 字符串 s 的所有无重复子串
// 思路：sliding window
// time O(N) space O(N)
func noDubSubstr(s string) (ans []string) {
	if len(s) == 0 {
		return
	}
	window := make(map[byte]int, 0)
	left, right := 0, 0
	for right < len(s) {
		c := s[right]
		window[c]++
		right++

		// 窗口内有重复的字符
		// 进行左缩，直到没有重复字符为止
		for window[c] > 1 {
			d := s[left]
			left++
			window[d]--
		}
		ans = append(ans, s[left:right])
	}
	return
}

// 438. Find All Anagrams in a String
// 438. 找到字符串中所有字母异位词
// 思路：滑动窗口
// time O(N) space O(N)
func findAnagrams(s string, p string) (ans []int) {
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return
	}
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)

	for i := 0; i < len(p); i++ {
		e := p[i]
		need[e]++
	}

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// 右扩窗口
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++

		// 根据条件，左缩窗口
		for right-left == len(p) {
			// 更新结果集
			if valid == len(need) {
				ans = append(ans, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return
}

// 567. Permutation in String
// 567. 字符串的排列
// 思路：sliding window (two pointer)
// time O(N) space O(N)
func checkInclusion(t string, s string) bool {
	if len(s) == 0 || len(s) < len(t) {
		return false
	}
	// need 记录字符串 t 中的字符出现的次数
	// window 记录窗口中对应 t 中的字符出现的次数
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)

	// 构建所需字符数的统计
	for i := 0; i < len(t); i++ {
		e := t[i]
		need[e]++
	}

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// 加入元素，更新统计，右扩窗口
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++

		// t 是 s 的子串，右扩长度 == len(t)时，才停止
		// 可确定左侧窗口需要收缩
		for right-left == len(t) {
			// need 中的所有字符统计数都满足时
			// 找到的子串长度等于 t && need 的字符对齐
			// 则找到对应全排列子串
			if valid == len(need) {
				return true
			}
			// 左缩窗口：移出元素，更新字符数据统计
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

// 76. Minimum Window Substring
// 76. 最小覆盖子串
// 思路：sliding window (two pointer)
// time O(N) space O(N)
func minWindow(s string, t string) string {
	if len(s) == 0 || len(s) < len(t) {
		return ""
	}
	// need 记录字符串 t 中的字符出现的次数
	// window 记录窗口中对应 t 中的字符出现的次数
	need, window := make(map[byte]int, 0), make(map[byte]int, 0)

	// 构建所需字符数的统计
	for i := 0; i < len(t); i++ {
		e := t[i]
		need[e]++
	}

	// 记录最小覆盖子串：起始索引、长度
	start, minLen := 0, math.MaxInt32

	left, right, valid := 0, 0, 0
	for right < len(s) {
		// 加入元素，更新统计，右扩窗口
		c := s[right]
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		right++

		// need 中的所有字符统计数都满足时
		// valid == len(need)
		for valid == len(need) {

			// 更新[最小覆盖子串]结果
			subLen := right - left
			if subLen < minLen {
				// 更新最小覆盖子串：起始索引，长度
				start = left
				minLen = subLen
			}

			// 左缩窗口：移出元素，更新字符数据统计
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}
