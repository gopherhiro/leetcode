package number

// 136. Single Number
// 136. 只出现一次的数字
// 思路：异或原理
// 1、a ^ 0 = a
// 2、a ^ a = 0
// 3、a ^ b ^ a = a ^ a ^ b = b
// time O(N) space O(1)
func singleNumber(nums []int) int {
	x := 0
	for _, v := range nums {
		x ^= v
	}
	return x
}
