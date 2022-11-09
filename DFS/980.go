package DFS

// 980. Unique Paths III
// 980. 不同路径 III
// 思路：DFS
func uniquePathsIII(grid [][]int) int {
	answer := 0
	m, n := len(grid), len(grid[0])

	// 1、找到起点
	// 2、统计走到终点需要走的步数
	// init empty == 1, 走完所有空格之后，需要再走一步
	// 才能到达终点，所以，走到终点所需步数为：空格数 + 1，1 预先设置。
	starti, startj := 0, 0
	empty := 1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				starti = i
				startj = j
				continue
			}
			if grid[i][j] == 0 {
				empty++
			}
		}
	}

	var dfs func(grid [][]int, i, j, step int)
	dfs = func(grid [][]int, i, j, step int) {
		// 超出索引范围，丢弃
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		// 障碍不能走
		if grid[i][j] == -1 {
			return
		}
		// 走过不能走
		if grid[i][j] == -2 {
			return
		}

		// 达到终点，就要停下来，不要瞎走了，很浪费体力的。
		// 能达到终点的路径可能有很多条，但只有走过了所有的空格的路径，才是我们需要的。
		// 如果走到终点后没有停下来，会导致永远搜索不到结果，只能等待超出索引范围结束。
		if grid[i][j] == 2 {
			if step == empty {
				answer++
			}
			return
		}

		// 标记走过
		grid[i][j] = -2

		dfs(grid, i-1, j, step+1) // move up
		dfs(grid, i+1, j, step+1) // move down
		dfs(grid, i, j-1, step+1) // move left
		dfs(grid, i, j+1, step+1) // move right

		// 从 [i,j] 为中心，向四个方向发散走，进行条件统计
		// 在统计完后需要回溯状态，以便于下一个点的统计。
		grid[i][j] = 0
	}

	dfs(grid, starti, startj, 0)
	return answer
}
