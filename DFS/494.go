package DFS

/*
with array nums = [1,1,1], targer = 3 example,
binary tree structure as follows:
                1
[0]		 +1		     -1
[1]	  +1    -1     +1   -1
[2] +1 -1 +1 -1  +1 -1 +1 -1
	 y  n  n  n   n  n  n  n
*/
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
