package number

// 191. Number of 1 Bits
// 191. 位1的个数
// 思路：验证num最后一位是不是1
// 1、是1，则+1
// 2、非1，则+0
// time O(N) space O(1)
func hammingWeightL(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	return count
}

// 191. Number of 1 Bits
// 191. 位1的个数
// 思路：取余法
// time O(N) space O(1)
func hammingWeightK(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num % 2)
		num = num >> 1
	}
	return count
}

// 191. Number of 1 Bits
// 191. 位1的个数
// 思路：bitwise
// n & (n-1) 作用是消除数字 n 的二进制表示中的最后一个 1。
// time O(N) space O(1)
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
