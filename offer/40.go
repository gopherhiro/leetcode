package offer

import "sort"

// 剑指 Offer 40. 最小的k个数
// 思路：排序
// time O(nlogn) space O(logn)
func getLeastNumbers(arr []int, k int) []int {
	if len(arr) == 0 {
		return []int{}
	}
	sort.Ints(arr)
	ans := make([]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, arr[i])
	}
	return ans
}
