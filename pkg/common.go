package pkg

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func GenerateMemo2D(row, col, initVal int) [][]int {
	memo := make([][]int, row)
	for i := range memo {
		memo[i] = make([]int, col)
		for j := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}

func GenerateMemoD(n, initVal int) []int {
	memo := make([]int, n)
	for i := range memo {
		memo[i] = initVal
	}
	return memo
}

func GenerateDp(row, col int) [][]int {
	dp := make([][]int, row)
	for i := range dp {
		dp[i] = make([]int, col)
	}
	return dp
}

func GenerateBoolDp(row, col int) [][]bool {
	dp := make([][]bool, row)
	for i := range dp {
		dp[i] = make([]bool, col)
	}
	return dp
}
