package island

// 695. Max Area of Island
// 695. 岛屿的最大面积
// 思路：Depth First Search
// 1、遍历所有岛屿并统计岛屿所拥有的陆地数量
// 2、择优出陆地数量最多的岛屿，返回其陆地数量
// time O(MN) space O(MN)
func maxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	const (
		LAND  = 1
		WATER = 0
	)
	m, n := len(grid), len(grid[0])

	var dfs func(grid [][]int, i, j int) int
	dfs = func(grid [][]int, i, j int) int {
		// 边界情况处理
		if i < 0 || j < 0 || i >= m || j >= n {
			return 0
		}
		if grid[i][j] == WATER {
			return 0
		}
		grid[i][j] = WATER
		up := dfs(grid, i-1, j)     // 上
		bottom := dfs(grid, i+1, j) // 下
		left := dfs(grid, i, j-1)   // 左
		right := dfs(grid, i, j+1)  // 右
		return 1 + up + bottom + left + right
	}

	maxS := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == LAND {
				S := dfs(grid, i, j)
				maxS = max(maxS, S)
			}
		}
	}
	return maxS
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
