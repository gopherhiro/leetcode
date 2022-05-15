package twosum

import "sort"

// 18. 4Sum
// 18. 四数之和
// 思路：sort + two pointer
// time O(N^2) space O(1)
func fourSum(nums []int, target int) [][]int {
	if len(nums) <= 1 {
		return [][]int{}
	}
	// sort
	sort.Ints(nums)
	return nSumTarget(nums, 4, 0, target)
}
