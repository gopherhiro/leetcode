package string

// 2446. Determine if Two Events Have Conflict
// 2446. 判断两个事件是否存在冲突
// 思路：字符串比较
// time O(1) space O(1)
func haveConflict(e1 []string, e2 []string) bool {
	// 1 2 2 1
	// 0 1 0 1
	return e1[0] <= e2[1] && e2[0] <= e1[1]
}

func haveConflict2(e1 []string, e2 []string) bool {
	return max(e1[0], e2[0]) <= min(e1[1], e2[1])
}

func max(x, y string) string {
	if x > y {
		return x
	}
	return y
}

func min(x, y string) string {
	if x < y {
		return x
	}
	return y
}
