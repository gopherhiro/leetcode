package backtracking

// 39. Combination Sum
// 39. 组合总和
// 思路：回溯+剪枝。
func combinationSum(candidates []int, target int) (res [][]int) {
	if len(candidates) == 0 {
		return
	}
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
			// 做出选择
			target -= candidates[i]
			track = append(track, candidates[i])

			// 进入下一层决策树
			// 往下一层传[i]，每一个数字都可以重复使用。
			backtrack(i)

			// 撤销选择
			track = track[:len(track)-1]
			target += candidates[i]
		}

	}
	backtrack(0)
	return
}
