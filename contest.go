package main

import "fmt"

func main() {
	gem := []int{0, 2, 5, 4}
	operations := [][]int{
		{3, 2},
		{3, 2},
		{1, 3},
		{0, 2},
		{3, 0},
		{3, 1},
		{0, 3},
		{2, 1},
		{3, 0},
	}
	r := giveGem(gem, operations)
	fmt.Println(r)
}

// LCP 50. 宝石补给
// 思路：场景模拟
// time O(N) space O(1)
// N 是 max(len(gem),len(operations))
func giveGem(gem []int, operations [][]int) int {
	if len(gem) == 0 {
		return 0
	}
	for i := 0; i < len(operations); i++ {
		x := operations[i][0]
		y := operations[i][1]

		// gem[x] 的一半
		give := gem[x] / 2

		gem[x] -= give // 授人玫瑰，减少
		gem[y] += give // 得到馈赠，增加
	}
	// 找寻最大、最小值，求差
	mi, ma := gem[0], gem[0]
	for i := 0; i < len(gem); i++ {
		if gem[i] < mi {
			mi = gem[i]
		}
		if gem[i] > ma {
			ma = gem[i]
		}
	}
	return ma - mi
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
