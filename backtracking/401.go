package backtracking

import "fmt"

// 401. Binary Watch
// 401. 二进制手表
// 思路：此题是一个不错的回溯算法思想运用题。
func readBinaryWatch(turnedOn int) (res []string) {
	if turnedOn < 0 {
		return
	}

	// 可选列表：LED灯，前4个为小时，后6个为分钟。
	times := []int{8, 4, 2, 1, 32, 16, 8, 4, 2, 1}

	var backtrack func(turnedOn, start, hours, minutes int)
	backtrack = func(turnedOn, start, hours, minutes int) {
		// 结束条件
		if turnedOn == 0 {
			// 剪枝：丢弃不合法时间。
			if hours > 11 || minutes > 59 {
				return
			}
			var tmpTime string

			// 处理小时
			tmpTime += fmt.Sprintf("%d", hours)
			tmpTime += ":"

			// 处理分钟
			if minutes < 10 {
				tmpTime += fmt.Sprintf("0%d", minutes)
			} else {
				tmpTime += fmt.Sprintf("%d", minutes)
			}
			// 收集时间结果
			res = append(res, tmpTime)
			return
		}

		// hours 需要加起来
		// minutes 需要加起来

		for i := start; i < len(times); i++ {
			// 做出选择
			isHour := i < 4
			if isHour {
				hours += times[i]
			} else {
				minutes += times[i]
			}

			// 进入下一层决策树
			// 进入下一层决策树，LED灯打开的数量，需要减1。
			backtrack(turnedOn-1, i+1, hours, minutes)

			// 撤销选择
			if isHour {
				hours -= times[i]
			} else {
				minutes -= times[i]
			}
		}

	}

	backtrack(turnedOn, 0, 0, 0)

	return
}
