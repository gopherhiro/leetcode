package main

import "fmt"

func main() {
	r := gen3Dp(2, 3, 4)
	fmt.Println(r)
}

func gen3Dp(x, y, z int) [][][]int {
	dp := make([][][]int, x)
	for i := range dp {
		dp[i] = make([][]int, y)
		for j := range dp[i] {
			dp[i][j] = make([]int, z)
		}
	}
	return dp
}
