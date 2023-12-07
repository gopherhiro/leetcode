package array

import "sort"

// 这道题的解法有许多种：
// 1. 位运算
// 2. 作差法
// 3. 数组哈希 or 哈希表
// 4. 排序 & 比较

// 268. Missing Number
// 268. 丢失的数字
// 思路：位运算：异或
// a ^ a = 0
// a ^ 0 = a
// time O(N) space O(1)
// 找缺失数、找出现一次数都是异或的经典应用。
// 我们可以先求得 [1,n] 的异或和 ans，然后用 ans 对各个 nums[i] 进行异或。
// 这样最终得到的异或和表达式中，只有缺失元素出现次数为 111 次，其余元素均出现两次（x⊕x=0），即最终答案 ans 为缺失元素。
func missingNumberB(nums []int) int {
	n := len(nums)

	ans := 0

	// 索引、元素做异或
	for i := 0; i <= n; i++ {
		ans ^= i
	}

	for _, v := range nums {
		ans ^= v
	}

	return ans
}

// 268. Missing Number
// 268. 丢失的数字
// 思路：求和算差，差值就是那个缺失的数。
// time O(N) space O(1)
// 推荐这种解法：比较好理解。
func missingNumber(nums []int) int {
	n := len(nums)
	sn := n * (n + 1) / 2

	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sn - sum
}

// 思路：一个简单的做法是直接对 nums 进行排序，
// 找到符合 nums[i]!=i 的位置即是答案，
// 如果不存在 nums[i]!=i 的位置，则 n 为答案。
// time O(nlogn) space O(logn)
func missingNumberC(nums []int) int {
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if nums[i] != i {
			return i
		}
	}
	return n
}
