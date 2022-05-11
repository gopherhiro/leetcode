package array

import "sort"

/********* 如何同时寻找缺失和重复的元素 ********/
// 645. Set Mismatch
// 645. 错误的集合
// 思路：排序+双指针
// 找寻重复元素：如果相邻的两个元素相等，则该元素为重复的数字。
// 找寻丢失元素：
// 1、当丢失的数字小于 n 时，计算当前元素与上一个元素的差 > 1，即可找到丢失的数字.
// 2、如果 nums[n-1] != n，则丢失的数字是 n.
// time O(N*logN) space O(logN)
func findErrorNumsN(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	ans := make([]int, 2)

	sort.Ints(nums)

	prev := 0
	for _, curr := range nums {
		// 发现重复数字
		if curr == prev {
			ans[0] = curr
		}
		// 发现丢失数字
		if curr-prev > 1 {
			ans[1] = prev + 1
		}
		// prev 指向上一个元素
		prev = curr
	}
	// 排序后，nums[n-1] != n，表示丢失的数字是n
	// 特判一下：丢失的数字是 n
	if nums[n-1] != n {
		ans[1] = n
	}
	return ans
}

// 645. Set Mismatch
// 645. 错误的集合
// return [dup,lost]
// 思路：哈希表
// 1、遍历一遍数组，记录每个数字出现的次数。
// 2、遍历 [1...N]，找寻重复和丢失的元素。
// - 重复的元素出现2次，丢失的元素出现0次。
// time O(N) space O(N)
func findErrorNums(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	ans := make([]int, 2)

	// 遍历一遍数组，记录每个数字出现的次数
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
	}
	// 遍历 [1...N]，找寻重复和丢失的元素即可。
	// 重复的元素出现2次
	// 丢失的元素出现0次
	for i := 1; i <= n; i++ {
		count := hash[i]
		if count == 2 {
			// 重复的数
			ans[0] = i
		}
		if count == 0 {
			// 丢失的数
			ans[1] = i
		}
	}
	return ans
}
