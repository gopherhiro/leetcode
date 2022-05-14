package binarysearch

// 1011. Capacity To Ship Packages Within D Days
// 1011. 在 D 天内送达包裹的能力
// 思路：二分搜索之左侧边界，寻找运载能力最小值。
// time O(N*logM), N 是 weights中的元素个数， M是其所有元素之和。
func shipWithinDays(weights []int, limitedDays int) int {
	if len(weights) == 0 || limitedDays == 0 {
		return 0
	}
	// 船的最小载重应该是weights数组中元素的最大值，因为每次至少得装一件货物走。
	// 最大载重显然就是weights数组所有元素之和，也就是一次把所有货物都装走。
	left, right := getMax(weights), getSum(weights)
	for left <= right {
		mid := left + (right-left)/2
		if canFinishShip(weights, mid, limitedDays) {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	return left
}

// 以 capacity 的运载能力，能否在限定的天数内完成运输
func canFinishShip(weights []int, capacity, days int) bool {
	// needDays 初始值最小为1(题目要求)。
	sum, needDays := 0, 1
	for _, weight := range weights {
		if sum+weight > capacity {
			needDays++
			sum = 0
		}
		sum += weight
	}
	return needDays <= days
}

func getSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
