package backtracking

import (
	"math"
	"sort"
)

// 47. Permutations
// 47. 全排列
// 思路：回溯
// time O(n*n!), space O(n)
func permuteUnique(nums []int) (ans [][]int) {
	// 先排序，让相同的元素在一起
	sort.Ints(nums)

	// 记录路径
	track := make([]int, 0)

	// 空间换时间：使用map辅助，剪枝的时间复杂度可降为O(1)
	used := make(map[int]bool)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for i, num := range nums {
			// 剪枝，排除不合理的情况。
			// 该下标元素是否使用过。
			if used[i] {
				continue
			}
			// 新增额外剪枝：把重复元素的排列去掉。
			// 固定相同的元素在排列中的相对位置：
			// 即如果前面的相邻 相等元素没有用过，则剪枝。
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}

			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)
			used[i] = true

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}

	}

	backtrack(nums)

	return
}

// 47. Permutations
// 47. 全排列
// 思路：回溯
// time O(n*n!), space O(n)
func permuteUniqueII(nums []int) (ans [][]int) {
	// 先排序，让相同的元素在一起
	sort.Ints(nums)

	// 记录路径
	track := make([]int, 0)

	// 空间换时间：使用map辅助，剪枝的时间复杂度可降为O(1)
	used := make(map[int]bool)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		preNumber := math.MinInt8
		for i, num := range nums {
			// 剪枝，排除不合理的情况。
			if used[i] {
				continue
			}

			// 新增额外剪枝：把重复元素的排列去掉。
			// 要选的元素是否和上一次选择过的相等。
			if num == preNumber {
				continue
			}

			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)
			used[i] = true

			// 记录之前树枝上元素的值
			preNumber = num

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
			used[i] = false
		}

	}

	backtrack(nums)

	return
}
