package greedy

import "sort"

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
