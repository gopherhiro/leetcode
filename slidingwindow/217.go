package slidingwindow

import "sort"

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
