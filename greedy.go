package main

import (
	"fmt"
	"sort"
)

// 动态规划和贪心算法的关系。
// 贪心算法可以理解为一种特殊的动态规划问题。
// 拥有一些更特殊的性质，可以进一步降低动态规划算法的时间复杂度。

/**
贪心是一种在每一步选择中都采取在当前状态下最好或最优（即最有利）的选择，从而达到结果是最优的算法。
什么是贪心选择性质呢，简单说就是：每一步都做出一个局部最优的选择，最终的结果就是全局最优。
这是一种特殊性质，其实只有一部分问题拥有这个性质。
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
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	r := canCompleteCircuitN(gas, cost)
	fmt.Println(r)
}

// 134. Gas Station
// 134. 加油站
// 策略：数学图像法
// 思路：思路：剩余油量最低点作为起点，能够保证油量累加和一直大于等于 0。
// time O(N) space O(1)
func canCompleteCircuitN(gas []int, cost []int) int {
	n := len(gas)
	start := 0
	sum, minSum := 0, 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
		if sum < minSum {
			// 选择油量最低点作为起点
			// 最有可能保证环游一周
			start = i + 1
			minSum = sum
		}
	}
	if sum < 0 {
		// 总增加油量小于总消耗油量，无解
		return -1
	}
	if start == n {
		// 环形数组
		return 0
	}
	return start
}

// 134. Gas Station
// 134. 加油站
// 策略：贪心
// time O(N) space O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	sum := 0
	for i := 0; i < n; i++ {
		sum += gas[i] - cost[i]
	}
	if sum < 0 {
		// 总增加油量小于总消耗油量，无解
		return -1
	}

	start := 0
	remain := 0
	for i := 0; i < n; i++ {
		remain += gas[i] - cost[i]
		if remain < 0 {
			start = i + 1 // 更新起点
			remain = 0
		}
	}

	// 环形数组
	if start == n {
		return 0
	}
	return start
}

// 452. Minimum Number of Arrows to Burst Balloons
// 452. 用最少数量的箭引爆气球
// 策略：贪心
// 思路：寻找最多不重叠区间的个数
// 如果区间都是重叠的，那只需要一个箭头就可以穿透所有区间。
// 如果最多有 n 个不重叠的区间，那么就至少需要 n 个箭头穿透所有区间。
// 主要复杂度在于排序：time O(N*logN) space O(logN)
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	// 按照 end 从小到大排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	count := 1 // 至少有一个不重叠的区间
	xEnd := points[0][1]
	for _, point := range points {
		start := point[0]
		if start <= xEnd {
			// 排除重叠的区间
			continue
		}
		count++         // 找到下一个不重叠的区间，计数加1
		xEnd = point[1] // 更新 x
	}
	return count
}

// 435. Non-overlapping Intervals
// 435. 无重叠区间
// 策略：贪心
// 思路：
// 1、从区间集合 intervals 中选择一个区间 x，这个 x 是在当前所有区间中结束最早的（end 最小）。
// 2、把所有与 x 区间相交的区间从区间集合 intervals 中删除。
// 3、重复步骤 1 和 2，直到 intervals 为空为止。之前选出的那些 x 就是最大的不重叠子集的个数。
// 4、已经计算出「最大的不重叠子集的个数」，至少要「移除区间」的数量 = len(nums) - 「最大的不重叠子集的个数」。
// time O(N*logN) space O(logN)，排序所需要的时间空间复杂度。
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 按照 end 从小到大排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 1              // 至少有一个区间不重叠
	xEnd := intervals[0][1] // 排序后第一个区间就是 x
	for _, interval := range intervals {
		start := interval[0]
		if start < xEnd {
			// 排除重叠区间
			continue
		}
		count++            // 找到下一个不重叠区间，计数加 1
		xEnd = interval[1] // 更新 x
	}
	return len(intervals) - count
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
