package backtracking

// 1863. Sum of All Subset XOR Totals
// 1863. 找出所有子集的异或总和再求和
// 思路：
// 1、回溯求出所有的子集。
// 2、遍历所有子集，求异或并求和。
func subsetXORSumN(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {

		// 结束条件
		subset := make([]int, len(track))
		copy(subset, track)

		// 计算子集异或并求和
		if len(subset) != 0 {
			xor := 0
			for _, elem := range subset {
				xor ^= elem
			}
			sum += xor
		}

		// 遍历选择列表，进行路径选择
		for i := start; i < len(nums); i++ {
			// 做出选择
			track = append(track, nums[i])

			// 进入下一层决策树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, 0)

	return sum
}

// 1863. Sum of All Subset XOR Totals
// 1863. 找出所有子集的异或总和再求和
// 思路：
// 1、回溯求出所有的子集。
// 2、遍历所有子集，求异或并求和。
func subsetXORSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	subsets := make([][]int, 0)
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		// 结束条件
		tmp := make([]int, len(track))
		copy(tmp, track)
		subsets = append(subsets, tmp)

		// 遍历选择列表，进行路径选择
		for i := start; i < len(nums); i++ {
			// 做出选择
			track = append(track, nums[i])

			// 进入下一层决策树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}
	}
	backtrack(nums, 0)

	// 计算子集异或并求和
	if len(subsets) == 0 {
		return 0
	}
	sum := 0
	for _, subset := range subsets {
		if len(subset) == 0 {
			continue
		}
		xor := 0
		for _, elem := range subset {
			xor ^= elem
		}
		sum += xor
	}
	return sum
}
