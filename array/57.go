package array

// 57. Insert Interval
// 57. 插入区间
// 思路：三分法：前、中、后
// time O(n) space O(1)
func insert(intervals [][]int, new []int) [][]int {
	answer := make([][]int, 0)
	start, end := new[0], new[1]
	i := 0

	// add all the intervals ending before newInterval starts
	for i < len(intervals) && intervals[i][1] < start {
		answer = append(answer, intervals[i])
		i++
	}

	// merge all overlapping intervals to one considering newInterval
	for i < len(intervals) && intervals[i][0] <= end {
		start = min(intervals[i][0], start)
		end = max(intervals[i][1], end)
		i++
	}
	answer = append(answer, []int{start, end})

	// add the rest of all
	answer = append(answer, intervals[i:]...)
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 57. Insert Interval
// 57. 插入区间
// 思路：三分法。
// time O(n) space O(1)
func insert2(intervals [][]int, interval []int) [][]int {
	answer := make([][]int, 0)

	left, right := interval[0], interval[1]
	merged := false
	for _, v := range intervals {
		// 在「插入区间」左侧且无交集
		if v[1] < left {
			answer = append(answer, v)
			continue
		}

		// 在「插入区间」右侧且无交集
		if v[0] > right {
			// 之前计算的并集，如果还没有加入结果集，则加入。
			// 已加入，则忽略。
			if !merged {
				answer = append(answer, []int{left, right})
				merged = true
			}
			answer = append(answer, v)
			continue
		}

		// 与「插入区间」有交集，则计算它们的并集
		left = min(left, v[0])
		right = max(right, v[1])
	}
	// 特殊情况处理：
	// 情况一：
	// 如果「元素区间」与「插入区间」没有任何重叠
	// 则直接把「插入区间」添加到最后即可。
	// 情况二：
	// 如果「元素区间」包含「插入区间」，则需要把「并集」更新结果加入结果集
	if !merged {
		answer = append(answer, []int{left, right})
	}
	return answer
}
