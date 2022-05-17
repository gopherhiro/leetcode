package backtracking

import "sort"

// 90. Subsets II
// 90. 子集 II
// 思路：backtracking
// time O(N*2^N), space O(N)
func subsetsWithDup(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return
	}

	// 先排序，让相同的元素在一起。
	sort.Ints(nums)

	// 记录已选择路径
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		// 前序位置：收集结果。
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)

		// 使用 start 参数控制递归，进行剪枝。
		for i := start; i < len(nums); i++ {

			// 剪枝。值相同的相邻树枝，只遍历第一条。
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// 做选择
			track = append(track, nums[i])

			// 进入下一层抉择树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}
	backtrack(nums, 0)
	return
}
