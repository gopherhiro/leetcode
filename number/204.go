package number

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
