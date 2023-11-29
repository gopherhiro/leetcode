package backtracking

// 46. Permutations
// 46. 全排列
// 思路：回溯
// time O(n*n!), space O(n)
func permute(nums []int) (ans [][]int) {
	track := make([]int, 0)

	var backtrack func(nums []int)
	backtrack = func(nums []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, 0, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for _, num := range nums {
			// 剪枝，排除不合理的情况。
			if find(track, num) {
				continue
			}
			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

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

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
// time O(N)
func find(s []int, val int) bool {
	for _, item := range s {
		if item == val {
			return true
		}
	}
	return false
}

// 46. Permutations
// 46. 全排列
// 思路：回溯，优化了剪枝时的时间复杂度：O(N) -> O(1)
// time O(n*n!), space O(n)
func permuteMap(nums []int) (ans [][]int) {
	track := make([]int, 0)
	// 空间换时间：使用map辅助，剪枝的时间复杂度可降为O(1)
	used := make(map[int]bool)

	var backtrack func(nums []int, track []int)
	backtrack = func(nums []int, track []int) {
		// 结束条件：已选择路径与原始可选列表长度相等时
		// 即表示得到了一个全排列
		if len(track) == len(nums) {
			tmp := make([]int, len(track))
			copy(tmp, track)
			ans = append(ans, tmp)
			return
		}

		for _, num := range nums {
			// 剪枝，排除不合理的情况。
			if used[num] {
				continue
			}
			// 为了维护每个节点的选择列表和路径，需要：
			// 递归之前做出选择，
			// 递归之后撤销选择。

			// 做出选择
			track = append(track, num)
			used[num] = true

			// 进入下一层决策树
			backtrack(nums, track)

			// 撤销选择
			track = track[:len(track)-1]
			used[num] = false
		}

	}

	backtrack(nums, track)

	return
}
