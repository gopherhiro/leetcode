package main

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

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func generateDp(row, col int) [][]int {
	dp := make([][]int, row)
	for i, _ := range dp {
		dp[i] = make([]int, col)
	}
	return dp
}

func generateBoolDp(row, col int) [][]bool {
	dp := make([][]bool, row)
	for i, _ := range dp {
		dp[i] = make([]bool, col)
	}
	return dp
}

func generateMemo(n, initVal int) []int {
	memo := make([]int, n)
	for i := range memo {
		memo[i] = initVal
	}
	return memo
}
