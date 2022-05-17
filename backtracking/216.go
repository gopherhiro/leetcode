package backtracking

// 216. Combination Sum III
// 216. 组合总和 III
// 思路：回溯
func combinationSum3(k int, n int) (res [][]int) {
	if k <= 0 || n <= 0 {
		return
	}
	// 记录已选择路径
	track := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		// 结束条件
		if len(track) == k && n == 0 {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		if n < 0 {
			return
		}

		// 遍历选择列表，进行路径选择
		for i := start; i <= 9; i++ {
			// 剪枝：
			// track的长度加上区间 [start, n] 的长度小于 k
			// 不可能构造出长度为 k 的 track
			// 比如：start = n 时，
			if len(track)+(n-start)+1 < k {
				continue
			}

			// 做出选择
			n -= i
			track = append(track, i)

			// 进入下一层决策树
			backtrack(i + 1)

			// 撤销选择
			track = track[:len(track)-1]
			n += i
		}

	}
	backtrack(1)
	return
}
