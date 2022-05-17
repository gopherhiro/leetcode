package backtracking

import "sort"

// 40. Combination Sum II
// 40. 组合总和II
// 思路：回溯+剪枝。
func combinationSum2(candidates []int, target int) (res [][]int) {
	if len(candidates) == 0 {
		return
	}

	sort.Ints(candidates)

	// 记录已选择的路径
	track := make([]int, 0)
	var backtrack func(start int)
	backtrack = func(start int) {
		// 结束条件
		if target == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 不符合条件的都丢弃。
		if target < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			// 剪枝：值相同的树枝，只遍历第一条。
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// 做出选择
			target -= candidates[i]
			track = append(track, candidates[i])

			// 进入下一层决策树
			// 往下一层传[i + 1]，可以限定每个数字只能使用一次。
			backtrack(i + 1)

			// 撤销选择
			track = track[:len(track)-1]
			target += candidates[i]
		}

	}
	backtrack(0)
	return
}
