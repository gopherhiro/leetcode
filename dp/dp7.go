package main

import (
	"fmt"
	"math"
)

func main() {
	k, n := 3, 14
	r := superEggDrop(k, n)
	fmt.Println(r)
}

// 887. Super Egg Drop
// 887. 鸡蛋掉落
// 思路：dynamic programming
// dp 数组定义：dp[k][m] = n
// 有 k 个鸡蛋，最多可以尝试扔 m 次，能够确定F的最高楼层数。
// time O(KN) space O(KN)
func superEggDrop(k int, n int) int {
	if k == 0 || n == 0 {
		return 0
	}
	dp := generateDp(k, n)
	m := 0
	for dp[k][m] < n {
		m++
		for i := 1; i <= k; i++ {
			dp[i][m] = dp[i][m-1] + dp[i-1][m-1] + 1
		}
	}
	return m
}

func generateDp(m, n int) [][]int {
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}
	return dp
}

// 887. Super Egg Drop
// 887. 鸡蛋掉落
// 思路：recursion + memo
// dp 函数定义：最坏情况下的最少扔鸡蛋次数为 dp(k,n): k 个鸡蛋，测试 n 层楼。
// time O(K*N*logN) space O(KN)
func superEggDropMM(k int, n int) int {
	if k == 0 || n == 0 {
		return 0
	}

	memo := generateMemo(k, n, -1)

	var dp func(k, n int) int
	dp = func(k, n int) int {
		// base case
		// n == 0，不用扔鸡蛋
		if n == 0 {
			return 0
		}
		// k == 1，即只有一个鸡蛋，只能线性扫描所有层
		if k == 1 {
			return n
		}

		if memo[k][n] != -1 {
			return memo[k][n]
		}

		res := math.MaxInt32

		// 二分搜索代替线性搜索
		left, right := 1, n
		for left <= right {
			mid := left + (right-left)/2
			eggBroken := dp(k-1, mid-1) // 鸡蛋碎了 [1..mid-1]，共 mid-1 层。
			eggUnbroken := dp(k, n-mid) // 鸡蛋没碎 [mid+1..n]，共 n-mid 层。
			// 最坏情况下的最少扔鸡蛋次数
			num := max(eggBroken, eggUnbroken) + 1 // 在第 mid 层，扔了一次。
			res = min(res, num)
			if eggBroken > eggUnbroken {
				right = mid - 1 // 往低层找
			} else {
				left = mid + 1 // 往高层找
			}
		}
		memo[k][n] = res
		return res
	}
	return dp(k, n)
}

// 887. Super Egg Drop
// 887. 鸡蛋掉落
// 思路：recursion + memo
// dp 函数定义：最坏情况下的最少扔鸡蛋次数为 dp(k,n): k 个鸡蛋，测试 n 层楼。
// time O(k*n^2) space O(kn)
func superEggDropM(k int, n int) int {
	if k == 0 || n == 0 {
		return 0
	}

	memo := generateMemo(k, n, -1)

	var dp func(k, n int) int
	dp = func(k, n int) int {
		// base case
		// n == 0，不用扔鸡蛋
		if n == 0 {
			return 0
		}
		// k == 1，即只有一个鸡蛋，只能线性扫描所有层
		if k == 1 {
			return n
		}

		if memo[k][n] != -1 {
			return memo[k][n]
		}

		res := math.MaxInt32
		// for 循环遍历穷举所有选择
		for i := 1; i <= n; i++ {
			// 最坏情况下的最少扔鸡蛋次数
			eggBroken := dp(k-1, i-1)              // 鸡蛋碎了 [1..i-1]，共 i-1 层。
			eggUnbroken := dp(k, n-i)              // 鸡蛋没碎 [i+1..n]，共 n-i 层。
			num := max(eggBroken, eggUnbroken) + 1 // 在第 i 层，扔了一次。
			res = min(res, num)
		}
		memo[k][n] = res
		return res
	}
	return dp(k, n)
}

func generateMemo(m, n, initVal int) [][]int {
	memo := make([][]int, m+1)
	for i, _ := range memo {
		memo[i] = make([]int, n+1)
		for j, _ := range memo[i] {
			memo[i][j] = initVal
		}
	}
	return memo
}

// 887. Super Egg Drop
// 887. 鸡蛋掉落
// 思路：recursion
// dp 函数定义：最坏情况下的最少扔鸡蛋次数为 dp(k,n): k 个鸡蛋，测试 n 层楼。
// time O(k*n^2) space O(kn)
func superEggDropR(k int, n int) int {
	if k == 0 || n == 0 {
		return 0
	}
	var dp func(k, n int) int
	dp = func(k, n int) int {
		// base case
		// n == 0，不用扔鸡蛋
		if n == 0 {
			return 0
		}
		// k == 1，即只有一个鸡蛋，只能线性扫描所有层
		if k == 1 {
			return n
		}
		res := math.MaxInt32
		// for 循环遍历穷举所有选择
		for i := 1; i <= n; i++ {
			// 最坏情况下的最少扔鸡蛋次数
			eggBroken := dp(k-1, i-1)              // 鸡蛋碎了 [1..i-1]，共 i-1 层。
			eggUnbroken := dp(k, n-i)              // 鸡蛋没碎 [i+1..n]，共 n-i 层。
			num := max(eggBroken, eggUnbroken) + 1 // 在第 i 层，扔了一次。
			res = min(res, num)
		}
		return res
	}
	return dp(k, n)
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
