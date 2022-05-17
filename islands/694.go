package island

import (
	"fmt"
	"strings"
)

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
