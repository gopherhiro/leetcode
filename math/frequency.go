package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	nums := []int{1, 1, 1, 2, 3, 3, 3}
	o := Constructor(nums)
	r := o.Pick(1)
	fmt.Println(r)
}

/********* XX ********/



/********* 如何在无限序列中随机抽取元素 ********/
// 382. Linked List Random Node
// 382. 链表随机节点
// 思路：水塘抽样算法
// time O(N) space O(1)
type List struct {
	head *ListNode
}

func Constructor(head *ListNode) List {
	return List{
		head: head,
	}
}

func (l *List) GetRandom() int {
	var ans int
	i := 0
	p := l.head
	for p != nil {
		i++
		// 生成 [0,i) 之间整数
		// 该整数等于0的概率是 1/i
		if rand.Intn(i) == 0 {
			ans = p.Val
		}
		p = p.Next
	}
	return ans
}

// 398. Random Pick Index
// 398. 随机数索引
// 思路：水塘抽样算法
// time O(N) space O(1)
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (s *Solution) Pick(target int) int {
	ans, count := -1, 0
	for i, v := range s.nums {
		if v != target {
			continue
		}
		count++
		// 生成 [0,count) 之间整数
		// 该整数等于0的概率是 1/count
		if rand.Intn(count) == 0 {
			ans = i
		}
	}
	return ans
}

/********* 如何同时寻找缺失和重复的元素 ********/
// 645. Set Mismatch
// 645. 错误的集合
// 思路：排序+遍历
// 找寻重复元素：如果相邻的两个元素相等，则该元素为重复的数字。
// 找寻丢失元素：
// 1、当丢失的数字小于 n 时，计算当前元素与上一个元素的差 > 1，即可找到丢失的数字.
// 2、如果 nums[n-1] != n，则丢失的数字是 n.
// time O(N*logN) space O(logN)
func findErrorNumsN(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	ans := make([]int, 2)

	sort.Ints(nums)

	prev := 0
	for _, curr := range nums {
		// 发现重复数字
		if curr == prev {
			ans[0] = curr
		}
		// 发现丢失数字
		if curr-prev > 1 {
			ans[1] = prev + 1
		}
		// prev 指向上一个元素
		prev = curr
	}
	// 丢失的数字是 n，特判一下
	if nums[n-1] != n {
		ans[1] = n
	}
	return ans
}

// 645. Set Mismatch
// 645. 错误的集合
// return [dup,lost]
// 思路：哈希表
// 1、遍历一遍数组，记录每个数字出现的次数。
// 2、遍历 [1...N]，找寻重复和丢失的元素即可。
// time O(N) space O(N)
func findErrorNums(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	ans := make([]int, 2)

	// 遍历一遍数组，记录每个数字出现的次数
	hash := make(map[int]int, 0)
	for _, v := range nums {
		hash[v]++
	}
	// 遍历 [1...N]，找寻重复和丢失的元素即可。
	// 重复的元素出现2次
	// 丢失的元素出现0次
	for i := 1; i <= n; i++ {
		count := hash[i]
		if count == 2 {
			// 重复的数
			ans[0] = i
		}
		if count == 0 {
			// 丢失的数
			ans[1] = i
		}
	}
	return ans
}

/********* 如何高效进行模幂运算 ********/
// 372. Super Pow
// 372. 超级次方
// 思路：(a*b) mod m == [(a mod m) * (b mod m)] mod m
// time O(N) space O(N)
func superPow(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}
	last := b[n-1]
	b = b[:n-1]
	partOne := pow(a, last)
	partTwo := pow(superPow(a, b), 10)
	return partOne * partTwo % 1337
}

// 累乘求幂法
// time O(k) space O(1)
func pow(a, k int) int {
	ans := 1
	for i := 0; i < k; i++ {
		ans *= a
		ans %= 1337
	}
	return ans % 1337
}

// 求pow
// 求次方
// 思路：高效求幂法
// b 为奇数时：a^b = a * a^(b-1)
// b 为偶数时：a^b = [a ^ (b/2)] ^ 2
// time O(logK) space O(logK)
func myPow(a, k int) int {
	if k == 0 {
		return 1
	}
	// k 是奇数
	if k%2 == 1 {
		return a * myPow(a, k-1)
	}
	// k 是偶数
	sub := myPow(a, k/2)
	return sub * sub
}

/********* 如何高效寻找素数 ********/
// 204. Count Primes
// 204. 计数质数（找出 [2,n)范围的质数个数）
// 思路：埃拉托斯特尼筛法
// time O(N*loglogN) space O(N)
func countPrimes(n int) int {
	if n <= 1 {
		return 0
	}
	count := 0
	prime := prime(n)
	for i := 2; i < n; i++ {
		if prime[i] {
			count++
			// 提前标记非质数：i 是质数，则 i 的倍数不是质数。
			// j 从 i*i 开始，相比从 2*i 开始，可以跳过冗余计算。
			for j := i * i; j < n; j += i {
				prime[j] = false
			}
		}
	}
	return count
}

func prime(n int) []bool {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	return isPrime
}

// 204. Count Primes
// 204. 计数质数（找出 [2,n)范围的质数个数）
// 思路：质数定义
// time O(N*sqrtN) space O(1)
func countPrimesN(n int) int {
	if n <= 1 {
		return 0
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrimeN(i) {
			count++
		}
	}
	return count
}

// 12 = 2 × 6
// 12 = 3 × 4
// 12 = sqrt(12) × sqrt(12)
// 12 = 4 × 3
// 12 = 6 × 2
// 存在计算冗余，i 不需要遍历到 n，而只需要到 sqrt(n) 即可。
func isPrimeN(n int) bool {
	for i := 2; i*i <= n; i++ {
		// n 可以被除了 1 和它本身之外的数除尽，则不是质数。
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 204. Count Primes
// 204. 计数质数（找出 [2,n)范围的质数个数）
// 思路：质数定义
// time O(N^2) space O(1)
// 超时
func countPrimesM(n int) int {
	if n <= 1 {
		return 0
	}
	count := 0
	for i := 2; i < n; i++ {
		if isPrimeM(i) {
			count++
		}
	}
	return count
}

func isPrimeM(n int) bool {
	for i := 2; i < n; i++ {
		// n 可以被除了 1 和它本身之外的数除尽，则不是质数。
		if n%i == 0 {
			return false
		}
	}
	return true
}
