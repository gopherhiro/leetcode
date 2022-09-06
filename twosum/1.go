package twosum

import "sort"

// 1. 两数之和
// 1. Two Sum
// 思路：hash + 遍历
// time O(N), space O(N)
func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for k, v := range nums {
		if i, ok := hash[target-v]; ok {
			return []int{i, k}
		}
		hash[v] = k
	}
	return []int{-1, -1}
}

// 1. 两数之和
// 1. Two Sum
// 思路：排序 + 双指针
func twoSumP(nums []int, target int) []int {
	sort.Ints(nums)
	var helper func(nums []int, target int) []int
	helper = func(nums []int, target int) []int {
		left, right := 0, len(nums)-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				return []int{left + 1, right + 1}
			}
			if sum < target {
				left++
			}
			if sum > target {
				right--
			}
		}
		return []int{-1, -1}
	}
	return helper(nums, target)
}
