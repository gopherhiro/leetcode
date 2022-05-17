package backtracking

// 77. Combinations
// 77. 组合
// 思路：backtracking
// time O(N * 2^N) , space O(N)
func combine(n int, k int) (res [][]int) {
	if n <= 0 || k <= 0 {
		return
	}
	// 记录已选择的路径
	track := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(track) == k {
			tmp := make([]int, len(track))
			copy(tmp, track)
			res = append(res, tmp)
			return
		}

		// 使用 start 参数控制递归，进行剪枝。
		for i := start; i <= n; i++ {
			// 剪枝：
			// track的长度加上区间 [start, n] 的长度小于 k
			// 不可能构造出长度为 k 的 track
			// 比如：start = n 时，
			if len(track)+(n-start+1) < k {
				return
			}

			// 做选择
			track = append(track, i)

			// 进入下一层决策树
			backtrack(i + 1)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}
	// n >= 1
	backtrack(1)
	return
}
