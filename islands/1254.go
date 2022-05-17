package island

// 1254. Number of Closed Islands
// 1254. 统计封闭岛屿的数目
// 理解：封闭岛屿，即四周(left,right,up,bottom)全被海水包围的岛屿。
// 算法：Depth First Search
// 思路：把靠边的岛屿去掉（都变成海水），则剩下的岛屿都是封闭的了，再统计岛屿即可。
// 1：海水
// 0：陆地
// time O(MN) space O(MN)
func closedIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	const (
		LAND  = 0
		WATER = 1
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

	// 统计剩下的岛屿，即是封闭岛屿
	answer := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == LAND {
				answer++
				dfs(grid, i, j)
			}
		}
	}
	return answer
}
