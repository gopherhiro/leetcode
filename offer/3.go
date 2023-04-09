package offer

import "sort"

// 剑指 Offer 03. 数组中重复的数字
// 思路：哈希表
// time O(N) space O(N)
func findRepeatNumber(nums []int) int {
	has := make(map[int]bool, len(nums))
	for _, num := range nums {
		// 存在即返回
		if has[num] {
			return num
		}
		has[num] = true
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	ht := make(map[int]int, 0)
	for _, v := range nums {
		ht[v]++
		// 出现2次，表示重复
		if ht[v] == 2 {
			return v
		}
	}
	return -1
}

// 剑指 Offer 03. 数组中重复的数字
// 思路：排序
// time O(NlogN) space O(logN)
func findRepeatNumberS(nums []int) int {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 排序后，前后两个元素相同，发现重复。
		if nums[i] == nums[i+1] {
			return nums[i]
		}
	}
	return -1
}

// 剑指 Offer 03. 数组中重复的数字
// 思路：原地交换, element value in [0,n-1]
// 数组元素的 索引 和 值 是 一对多 的关系。
// 可遍历数组并通过交换操作，使元素的「索引与值」一一对应（nums[i] == i），
// 因而，就能通过索引映射对应的值，起到与字典等价的作用。
// time O(N) space O(1)
func findRepeatNumberE(nums []int) int {
	i, n := 0, len(nums)
	for i < n {
		// nums[i] == i, 索引与索引位的值对应，即人才匹配。
		if nums[i] == i {
			i++
			continue
		}
		// 找到重复的值，即溢出的人才。
		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		// swap 人才位，进行匹配。
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
