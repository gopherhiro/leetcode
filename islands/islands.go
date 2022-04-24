package main

import (
	"fmt"
	"strings"
)

func main() {
	grid := [][]int{
		{1, 1, 1, 0, 0},
		{0, 1, 1, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 1, 1},
	}
	r := numDistinctIslands(grid)
	fmt.Println(r)
}

// 694. Number of Distinct Islands
// 694. 不同岛屿的数量
// 思路：Depth First Search + Hashtable
// 方案：序列化岛屿（「自定义序列化路径代号」or 「陆地相对下标计算」）
// 该解法使用的是「自定义序列化路径代号」
// 1、通过「序列化方式」呈现「岛屿形状」，岛屿形状相同：遍历岛屿所形成的路径序列一样。
// 2、使用 hash table 存储去重，即得「不同岛屿的数量」。
// time O(MN) space O(MN)
func numDistinctIslands(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	const (
		LAND  = 1
		WATER = 0
	)

	const (
		start  = 0
		up     = 1
		bottom = 2
		left   = 3
		right  = 4
	)

	m, n := len(grid), len(grid[0])

	var dfs func(grid [][]int, i, j int, sb *strings.Builder, dir int)
	dfs = func(grid [][]int, i, j int, sb *strings.Builder, dir int) {
		// 边界情况处理
		if i < 0 || j < 0 || i >= m || j >= n {
			return
		}
		if grid[i][j] == WATER {
			return
		}
		grid[i][j] = WATER

		// 前序遍历位置：进入节点
		togo := fmt.Sprintf("%d,", dir)
		sb.WriteString(togo)

		dfs(grid, i-1, j, sb, up)     // 上
		dfs(grid, i+1, j, sb, bottom) // 下
		dfs(grid, i, j-1, sb, left)   // 左
		dfs(grid, i, j+1, sb, right)  // 右

		// 后序遍历位置：离开节点
		comeback := fmt.Sprintf("%d,", -dir)
		sb.WriteString(comeback)
	}

	hashtable := make(map[string]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == LAND {
				sb := &strings.Builder{}
				dfs(grid, i, j, sb, start)
				hashtable[sb.String()] += 1
			}
		}
	}
	return len(hashtable)
}

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
