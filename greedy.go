package main

import "fmt"

// 动态规划和贪心算法的关系。
// 贪心算法可以理解为一种特殊的动态规划问题。
// 拥有一些更特殊的性质，可以进一步降低动态规划算法的时间复杂度。

/**
贪心是一种在每一步选择中都采取在当前状态下最好或最优（即最有利）的选择，从而达到结果是最优的算法。
一个问题如果可以通过贪心法来解决，那么贪心法一般是解决这个问题的最好办法。
------------------------------------------------------
在所有的 feasible solution 中选择出 optimal solution
feasible solution（可行解）
optimal solution （最优解）
optimization problem （最优化问题）
dynamic programming, greedy algorithm is used for solving optimization problems

A greedy algorithm is an approach for solving a problem
by selecting the best option available at the moment.

Two properties:
1. Greedy Choice Property
If an optimal solution to the problem can be found
by choosing the best choice at each step without reconsidering the previous steps once chosen,
the problem can be solved using a greedy approach.This property is called greedy choice property.

2. Optimal Substructure
If the optimal overall solution to the problem corresponds to the optimal solution to its sub problems,
then the problem can be solved using a greedy approach. This property is called optimal substructure.
*/

func main() {
	nums := []int{2, 3, 1, 1, 4}
	r := jump(nums)
	fmt.Println(r)
}

// 45. Jump Game II
// 45. 跳跃游戏 II
// 策略：贪心
// 思路：每一次都选择跳得最远的位置，作为起跳点，即可保证最小的跳跃次数。
// time O(N) space O(1)
func jump(nums []int) int {
	jump, farthest, end := 0, 0, 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end {
			jump++         // 起跳
			end = farthest // 标记本次可以跳到的最远位置
		}
	}
	return jump
}

// 55. Jump Game
// 55. 跳跃游戏
// 思路：贪心思想
// 步骤：
// 1、每一步从当前位置 i 能够跳跃到的位置中，选择最远的那个跳远位置。
// 2、遍历整个数组，得到最远的跳跃长度，能够超越最后一个下标，则能够到达最后一个下标。
// - 在遍历的过程中，如果发现当前位置 i，最远的跳远位置是 i 本身，则表示从当前位置永远无法到达下一个下标，
// - 也就无法达到最后一个下标。
// time O(N) space O(1)
func canJump(nums []int) bool {
	farthest := 0
	n := len(nums)
	for i := 0; i < n-1; i++ {
		curCanJumpFarthest := i + nums[i] // 从当前位置 i 能够跳到的最远的位置
		farthest = max(farthest, curCanJumpFarthest)
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
