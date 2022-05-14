package math

import (
	"sort"
)

func main() {

}

// 231. Power of Two
// 231. 2 的幂
// 思路：bitwise operations
// 如果一个数是 2 的指数，那么它的二进制表示一定只含有一个 1。
// time O(1) space O(1)
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return n&(n-1) == 0
}

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

// 349. Intersection of Two Arrays
// 349. 两个数组的交集
// 思路：two set intersection
// time O(m+n) space O(m+n)
func intersection(nums1 []int, nums2 []int) (ans []int) {
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	s1, s2 := make(map[int]int, 0), make(map[int]int, 0)
	for _, num := range nums1 {
		s1[num]++
	}
	for _, num := range nums2 {
		s2[num]++
	}
	// 总是遍历较小的那个set
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for key, _ := range s1 {
		if _, ok := s2[key]; ok {
			ans = append(ans, key)
		}
	}
	return
}

// 66. Plus One
// 66. 加一
// 思路：找出最长的后缀 9
// time O(N) space O(1)
// 这个写法更加OK一些
func plusOneN(digits []int) []int {
	n := len(digits)
	// 从末尾开始，将 9 全部置零。
	// 找到第一个不为 9 的元素，将该元素加一即可。
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i] += 1
			return digits
		}
		digits[i] = 0
	}

	// 所有的元素均为 9
	// 构造一个长度比 digits 多 1 的新数组，
	// 将首元素置为 1，其余元素置为 0 即可。
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// 66. Plus One
// 66. 加一
// 思路：找出最长的后缀 9
// time O(N) space O(1)
func plusOne(digits []int) []int {
	n := len(digits)
	// 找出从末尾开始的第一个不为 9 的元素
	// 将该元素加一
	// 随后将末尾的 9 全部置零
	for i := n - 1; i >= 0; i-- {
		if digits[i] == 9 {
			continue
		}
		digits[i] += 1
		for j := i + 1; j < n; j++ {
			digits[j] = 0
		}
		return digits
	}

	// 所有的元素均为 9
	// 构造一个长度比 digits 多 1 的新数组，
	// 将首元素置为 1，其余元素置为 0 即可。
	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := append(nums1, nums2...)
	if len(num) == 0 {
		return 0
	}
	sort.Ints(num)
	n := len(num)

	mid := (n - 1) / 2
	if n%2 == 0 {
		sum := num[mid+1] + num[mid]
		return float64(sum) / 2
	}
	//奇数
	return float64(num[mid])
}

// 287. Find the Duplicate Number
// 287. 寻找重复数
// 思路：Sort + 遍历查找
func findDuplicate(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// 思路：遍历 + 哈希表
func findDuplicateN(nums []int) int {
	hash := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			return num
		}
		hash[num] = true
	}
	return -1
}

// 思路：Array as HashMap (Iterative)
func findDuplicateM(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}
	return nums[0]
}

// 思路：Cycle Detection，快慢指针
func findDuplicateO(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	// 因为下面即将判断slow != fast
	// 所以，这里需要提前走，然后才能进入判断。
	slow = nums[slow]
	fast = nums[nums[fast]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// Find the "entrance" to the cycle.
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
