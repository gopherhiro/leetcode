package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 1, 1}
	target := 3
	r := findTargetSumWays(nums, target)
	fmt.Println(r)
}

// 494. Target Sum
// 494. 目标和
// 思路：DFS
// time O(n*2^n) space O(n)
func findTargetSumWays(nums []int, target int) int {
	sum, count := 0, 0
	var dfs func(i, sum int)
	dfs = func(i, sum int) {
		if i == len(nums) && sum == target {
			count++
			return
		}

		if i > len(nums)-1 {
			return
		}

		dfs(i+1, sum+nums[i])
		dfs(i+1, sum-nums[i])
	}
	dfs(0, sum)
	return count
}
