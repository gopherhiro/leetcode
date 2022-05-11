package array

// 268. Missing Number
// 268. 丢失的数字
// 思路：位运算：异或
// a ^ a = 0
// a ^ 0 = a
// time O(N) space O(1)
func missingNumberB(nums []int) int {
	n := len(nums)

	ans := 0

	// 索引、元素做异或
	for i := 0; i < n; i++ {
		ans ^= i ^ nums[i]
	}
	// 最后的索引，起着补全索引对的作用
	ans ^= n

	return ans
}

// 268. Missing Number
// 268. 丢失的数字
// 思路：求和算差，差值就是那个缺失的数。
// time O(N) space O(1)
// 推荐这种解法：比较好理解。
func missingNumber(nums []int) int {
	n := len(nums)
	// 等差数列求和
	Sn := n * (n + 1) / 2

	// 数组内元素求和
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return Sn - sum
}
