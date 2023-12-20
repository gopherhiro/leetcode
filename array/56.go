package array

import "sort"

// 56. Merge Intervals
// 56. 合并区间
// 思路：sorting
// time O(nlogn) space O(logn)
func merge(intervals [][]int) [][]int {
	// sort the intervals by their start value
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, interval := range intervals {
		if len(merged) == 0 {
			merged = append(merged, interval)
			continue
		}
		last := merged[len(merged)-1]
		if last[1] < interval[0] {
			merged = append(merged, interval)
			continue
		}
		last[1] = max(last[1], interval[1])
	}
	return merged
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
