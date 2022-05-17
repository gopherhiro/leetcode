package island

// 1905. Count Sub Islands
// 1905. 统计子岛屿
// 思路：Depth First Search
// 1、当「岛屿B」的所有陆地在「岛屿A」中也是陆地时，则「岛屿B」是「岛屿A」的子岛屿。
// 2、如果「岛屿B」中存在一块陆地，在「岛屿A」中对应的位置是海水，那么「岛屿B」不是「岛屿A」的子岛屿。
// 方案：遍历 grid2 中的所有岛屿，把那些不可能是「子岛屿」的岛屿排除掉，剩下的就是「子岛屿」。
// time O(MN) space O(MN)
func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	if len(grid1) == 0 || len(grid2) == 0 {
		return 0
	}
	const (
		LAND  = 1
		WATER = 0
	)
	m, n := len(grid1), len(grid1[0])

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

	// 淹没非「子岛屿」的岛屿
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid1[i][j] == WATER && grid2[i][j] == LAND {
				dfs(grid2, i, j)
			}
		}
	}

	// 统计剩余岛屿的数量，即是「子岛屿」的数量
	answer := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == LAND {
				answer++
				dfs(grid2, i, j)
			}
		}
	}
	return answer
}
