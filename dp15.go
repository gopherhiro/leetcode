package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [][]int{
		{0, -3},
	}
	r := calculateMinimumHPM(nums)
	fmt.Println(r)
}

// 174. 地下城游戏
// 174. Dungeon Game
// 思路：dynamic programming compress state
// dp 数组定义：dp[i][j] = x
// 骑士从 dungeon[i][j] 到达终点（右下角），所需的最低初始健康点数为 x
// time O(MN) space O(N)
func calculateMinimumHPM(g [][]int) int {
	m, n := len(g), len(g[0])
	dp := make([]int, n+1)
	// 状态转移
	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			// 边界处理
			if i == m || j == n {
				dp[j] = math.MaxInt32
				continue
			}
			// base case
			if i == m-1 && j == n-1 {
				dp[j] = 1
				if g[i][j] < 0 {
					dp[j] = -g[i][j] + 1
				}
				continue
			}
			dp[j] = min(dp[j], dp[j+1]) - g[i][j]
			if dp[j] <= 0 {
				dp[j] = 1
			}
		}
	}
	return dp[0]
}

// 174. 地下城游戏
// 174. Dungeon Game
// 思路：dynamic programming
// dp 数组定义：dp[i][j] = x
// 骑士从 dungeon[i][j] 到达终点（右下角），所需的最低初始健康点数为 x
// time O(MN) space O(MN)
func calculateMinimumHP(g [][]int) int {
	m, n := len(g), len(g[0])
	dp := genDp(m+1, n+1)

	// base case : dp[m-1][n-1]
	dp[m-1][n-1] = 1
	if g[m-1][n-1] < 0 {
		dp[m-1][n-1] = -g[m-1][n-1] + 1
	}

	// 状态转移
	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			// 边界处理
			if i == m || j == n {
				dp[i][j] = math.MaxInt32
				continue
			}
			// 跳过 base case
			if i == m-1 && j == n-1 {
				continue
			}
			down := dp[i+1][j]
			right := dp[i][j+1]
			ret := min(down, right) - g[i][j]
			if ret <= 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = ret
			}
		}
	}
	return dp[0][0]
}

func genDp(row, col int) [][]int {
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	return dp
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 174. 地下城游戏
// 174. Dungeon Game
// 思路：dynamic programming
// dp 函数定义：dp(i,j) = x
// 骑士从 dungeon[i][j] 到达终点（右下角），所需的最低初始健康点数为 x
// time O(MN) space O(MN)
func calculateMinimumHPR(g [][]int) int {
	m, n := len(g), len(g[0])

	memo := generateMemo(m, n, -1)

	/* 定义：从 (i, j) 到达右下角，所需最低初始健康点数是 dp(i,j) */
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == m-1 && j == n-1 {
			if g[i][j] >= 0 {
				return 1
			}
			return -g[i][j] + 1
		}
		// 边界情况处理:因为要取最小值，返回最大值，方便取最小值
		if i == m || j == n {
			return math.MaxInt32
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		// 状态转移
		ret := min(dp(i+1, j), dp(i, j+1)) - g[i][j]
		if ret <= 0 {
			memo[i][j] = 1
		} else {
			memo[i][j] = ret
		}
		return memo[i][j]
	}
	// 计算左上角到右下角所需的最小生命值
	return dp(0, 0)
}

func generateMemo(row, col, initVal int) [][]int {
	memo := make([][]int, row)
	for i := range memo {
		memo[i] = make([]int, col)
		for j := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}
