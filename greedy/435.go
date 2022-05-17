package greedy

import "sort"

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
