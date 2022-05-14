package binarysearch

// 875. Koko Eating Bananas
// 875. 爱吃香蕉的珂珂
// 思路：二叉搜索之左侧边界搜索。
func minEatingSpeed(piles []int, limitHour int) int {
	if len(piles) == 0 || limitHour == 0 {
		return 0
	}
	// left,right 表示吃香蕉的速度范围
	left, right := 1, getMax(piles)
	for left <= right {
		mid := left + (right-left)/2
		// 以 mid 速度能否在 limitHour 小时内吃完香蕉
		if canFinish(piles, mid, limitHour) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// 以  speed 速度，能否在 limitHour 小时内吃完所有香蕉
func canFinish(piles []int, speed, limitHour int) bool {
	sumHour := 0
	for _, number := range piles {
		sumHour += timeOf(number, speed)
	}
	return sumHour <= limitHour
}

// 以 speed 速度吃 number 个香蕉，需要的时间
func timeOf(number, speed int) (needTime int) {
	needTime = number / speed
	// 如果一小时内吃不完，则只能下一个小时再吃。
	// 所以，一次吃不完的，还需要加1个小时来吃完。
	if number%speed > 0 {
		needTime += 1
	}
	return
}

// 计算数组的最大值
func getMax(nums []int) int {
	max := 0
	for _, v := range nums {
		max = maxvalue(max, v)
	}
	return max
}

func maxvalue(x, y int) int {
	if x > y {
		return x
	}
	return y
}
