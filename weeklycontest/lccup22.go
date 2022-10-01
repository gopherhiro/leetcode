package main

import "fmt"

func main() {
	A := []int{5, 10, 16, -6, 15, 11, 3}
	B := []int{16, 22, 23, 23, 25, 3, -16}
	r := temperatureTrend(A, B)
	fmt.Println(r)
}

// 1. 气温变化趋势
// 对于第 i ~ (i+1) 天的气温变化趋势，将根据以下规则判断：
// 若第 i+1 天的气温 高于 第 i 天，为 上升 趋势
// 若第 i+1 天的气温 等于 第 i 天，为 平稳 趋势
// 若第 i+1 天的气温 低于 第 i 天，为 下降 趋势
// 已知 temperatureA[i] 和 temperatureB[i] 分别表示第 i 天两地区的气温。
// 组委会希望找到一段天数尽可能多，且两地气温变化趋势相同的时间举办嘉年华活动。
// 请分析并返回两地气温变化趋势相同的最大连续天数。

func temperatureTrend(A []int, B []int) int {
	ra, rb := make([]int, 0), make([]int, 0)
	length := len(A)
	for i := 1; i < length; i++ {
		a := compare(A[i-1], A[i])
		ra = append(ra, a)
		b := compare(B[i-1], B[i])
		rb = append(rb, b)
	}

	res, seqLen := 0, 0
	for i := 0; i < length-1; i++ {
		if ra[i] == rb[i] {
			seqLen++
			continue
		} else {
			res = max(res, seqLen)
			seqLen = 0
		}
	}
	res = max(res, seqLen)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func compare(before, after int) int {
	if after == before {
		return 0
	}
	if after > before {
		return 1
	}
	return -1
}
