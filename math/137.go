package math

// 137. Single Number II
// 137. 只出现一次的数字 II
// 思路：bitwise
// time O(N) space O(N)
// 备注：不是很理解这个，启发一下。
func singleNumberIIO(nums []int) int {
	x1, x2, mask := 0, 0, 0
	for _, v := range nums {
		x2 ^= x1 & v
		x1 ^= v
		mask = ^(x1 & x2) // ^ 也表示取反的意思
		x2 &= mask
		x1 &= mask
	}
	return x1
}

// 137. Single Number II
// 137. 只出现一次的数字 II
// 思路：hash table 统计次数
// time O(N) space O(N)
func singleNumberII(nums []int) int {
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
	}
	for elem, count := range hash {
		if count == 1 {
			return elem
		}
	}
	return -1
}
