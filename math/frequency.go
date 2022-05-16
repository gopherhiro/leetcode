package math

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{1, 1, 1, 2, 3, 3, 3}
	o := ConstructorA(nums)
	r := o.Pick(1)
	fmt.Println(r)
}

/********* XX ********/

type ListNode struct {
	Val  int
	Next *ListNode
}

/********* 如何在无限序列中随机抽取元素 ********/
// 382. Linked List Random Node
// 382. 链表随机节点
// 思路：水塘抽样算法
// time O(N) space O(1)
type List struct {
	head *ListNode
}

func ConstructorB(head *ListNode) List {
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

func ConstructorA(nums []int) Solution {
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
