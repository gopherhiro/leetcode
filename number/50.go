package number

// 50. Pow(x, n)
// 思路：迭代 + 快速幂
// time O(logn) space O(1)
func myPow(x float64, n int) float64 {
	if n >= 0 {
		return iterFastPow(x, n)
	}
	// n is negative
	return 1.0 / iterFastPow(x, -n)
}

func iterFastPow(x float64, n int) float64 {
	ans := 1.0
	for n > 0 {
		if n&1 == 1 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}

// 50. Pow(x, n)
// 思路：递归 + 快速幂
// time O(logn) space O(logn)
func myPowR(x float64, n int) float64 {
	if n >= 0 {
		return fastPow(x, n)
	}
	// n is negative
	return 1.0 / fastPow(x, -n)
}

func fastPow(x float64, n int) float64 {
	// base case
	if n == 0 {
		return 1
	}
	// n is even number
	if n%2 == 0 {
		return fastPow(x*x, n/2)
	}
	// n is odd number
	return x * fastPow(x*x, n/2)
}
