package twosum

import "sort"

// 15. 3Sum
// 15. 三数之和
// 思路：排序 + two pointer
// time O(N^2) space O(N)
func threeSum(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{}
	}
	// sort
	sort.Ints(nums)
	return nSumTarget(nums, 3, 0, 0)
}
