package backtracking

// 78. Subsets
// 78. 子集
// 思路：回溯
// time O(N*2^N), space O(N)
func subsets(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return
	}
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

// 列举所有全排列，同一个位置的数可以使用多次。
// time O(n*n!), space O(n)
func permuteBase(nums []int) (ans [][]int) {
	track := make([]int, 0)

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

		for _, num := range nums {
			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，递归之后撤销选择，即回溯。
			// 做出选择
			track = append(track, num)

			// 进入下一层决策树
			backtrack(nums)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}

	backtrack(nums)

	return
}
