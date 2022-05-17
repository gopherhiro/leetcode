package island

// 200. Number of Islands
// 200. 岛屿数量
// 思路：Depth First Search
// 1、相邻的陆地，共同形成一个岛屿。
// 2、使用 DFS 巡视「根据地」相邻陆地，收复相邻的陆地，形成一个岛屿。
// 3、DFS结束，表示已经走到海岸线回来，陆地收复完毕，岛屿统一！
// time O(MN) space O(MN)
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var dfs func(grid [][]byte, i, j int)
	dfs = func(grid [][]byte, i, j int) {
		// 边界情况处理
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		// 如果已经收复，则不用收复了。
		if grid[i][j] == '0' {
			return
		}
		// 标记已收复
		grid[i][j] = '0'
		// 巡视 && 收复相邻陆地
		dfs(grid, i-1, j) // 上
		dfs(grid, i+1, j) // 下
		dfs(grid, i, j-1) // 左
		dfs(grid, i, j+1) // 右
	}

	answer := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				answer++        // 遇到陆地，将其作为根据地，即新岛屿
				dfs(grid, i, j) // 收复其相邻陆地，形成统一岛屿
			}

		}
	}
	return answer
}
