package string

import "sort"

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
