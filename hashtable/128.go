package hashtable

import "sort"

// 128. Longest Consecutive Sequence
// 128. 最长连续序列
// 思路：hashtable
// time O(n) space O(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// hash table
	m := make(map[int]bool, 0)
	for _, v := range nums {
		m[v] = true
	}

	ans := 0
	for _, num := range nums {
		// find min start number to get the longest sequence
		// num 不是连续子序列的第一个，跳过
		if m[num-1] {
			continue
		}
		// num 是连续子序列的第一个，开始向上计算连续子序列的长度
		currNum := num
		currLen := 1
		for m[currNum+1] {
			currLen += 1
			currNum += 1
		}
		// 更新最长连续序列的长度
		ans = max(ans, currLen)
	}
	return ans
}

// 128. Longest Consecutive Sequence
// 128. 最长连续序列
// 思路：sorting
// time O(nlogn) space O(logn)
func longestConsecutiveS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	ans, count := 0, 1
	for i := 1; i < len(nums); i++ {
		// skip it and go the next number
		prev, curr := nums[i-1], nums[i]
		if prev == curr {
			continue
		}
		// prev + 1 == curr, get it
		if prev+1 == curr {
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
