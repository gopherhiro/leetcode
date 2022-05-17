package island

// 1020. Number of Enclaves
// 1020. 飞地的数量
// 算法技巧：Depth First Search
// 思路：
// 1、把靠边的岛屿去掉（都变成海水），则剩下的就是封闭的岛屿。
// 2、组成「封闭岛屿」的陆地，无法离开矩阵。
// 3、即统计组成「封闭岛屿」的陆地数量即是答案。
// 0：海水
// 1：陆地
// time O(MN) space O(MN)
func numEnclaves(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	const (
		LAND  = 1
		WATER = 0
	)
	m, n := len(grid), len(grid[0])

	var dfs func(grid [][]int, i, j int)
	dfs = func(grid [][]int, i, j int) {
		// 边界情况处理
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid[i][j] == WATER {
			return
		}
		grid[i][j] = WATER
		dfs(grid, i-1, j) // 上
		dfs(grid, i+1, j) // 下
		dfs(grid, i, j-1) // 左
		dfs(grid, i, j+1) // 右
	}
	// 淹没上下边缘岛屿
	for j := 0; j < n; j++ {
		dfs(grid, 0, j)   // 淹没上边缘岛屿
		dfs(grid, m-1, j) // 淹没下边缘岛屿
	}
	// 淹没左右边缘岛屿
	for i := 0; i < m; i++ {
		dfs(grid, i, 0)   // 淹没左边缘岛屿
		dfs(grid, i, n-1) // 淹没右边缘岛屿
	}

	// 统计「组成封闭岛屿」的陆地数量
	answer := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == LAND {
				answer++
			}
		}
	}
	return answer
}
