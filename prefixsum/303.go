package prefixsum

// 303. Range Sum Query - Immutable
// 303. 区域和检索 - 数组不可变
// 策略：前缀和
// time O(N), space O(N)
type NumArray struct {
	preSum []int
}

// 构造前缀和
func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	return NumArray{preSum: preSum}
}

// 计算索引 [left, right] 之间的 nums 元素的和
func (na *NumArray) SumRange(left int, right int) int {
	return na.preSum[right+1] - na.preSum[left]
}
