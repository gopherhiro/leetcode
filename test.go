package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(2)
	fmt.Println(r)

	r1 := rand.New()

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
