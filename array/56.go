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
		// If the current interval start after the previous interval end,
		// then they do not overlap and we can append the current interval to merged.
		// Otherwise, they do overlap, and we merge them by updating the end of the previous interval
		// if it is less than the end of the current interval.
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
