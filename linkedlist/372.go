package linkedlist

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
