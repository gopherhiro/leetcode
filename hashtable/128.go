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
		if m[num-1] {
			continue
		}
		// start
		count := 1
		for m[num+1] {
			count = count + 1
			num = num + 1
		}
		ans = max(ans, count)
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
		if nums[i] == nums[i-1] {
			continue
		}
		// prev + 1 == curr, get it
		if nums[i-1]+1 == nums[i] {
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
