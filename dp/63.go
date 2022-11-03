package dp

// 63. Unique Paths II
// 63. 不同路径 II
// 思路：DP
// time O(mn) space O(mn)
func uniquePathsWithObstacles(grid [][]int) int {
	const isObstacle int = 1
	// dp[i][j]:represents the number of paths from [0,0] to [i,j]
	m, n := len(grid), len(grid[0])
	dp := gendp(m, n)

	// 入口被堵住了，则直接 return 0
	if grid[0][0] == 1 {
		return 0
	}
	// base case:
	// 第一列，如果一个元素是障碍物，则其后续的元素都不可到达。
	for i := 0; i < m; i++ {
		if grid[i][0] == isObstacle {
			for i < m {
				dp[i][0] = 0
				i++
			}
		} else {
			dp[i][0] = 1
		}
	}
	// 第一行，如果一个元素是障碍物，则其后续的元素都不可到达。
	for j := 0; j < n; j++ {
		if grid[0][j] == isObstacle {
			for j < n {
				dp[0][j] = 0
				j++
			}
		} else {
			dp[0][j] = 1
		}
	}

	//  grid[i][j] == 1 代表此路不通，则 dp[i][j] = 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == isObstacle {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func gendp(m, n int) [][]int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	return dp
}
