package number

// 202. Happy Number
// 202. 快乐数
// 思路：哈希表检测循环
// time O(logn) space O(logn)
func isHappy(n int) bool {
	seen := make(map[int]bool, 0)
	for n != 1 && !seen[n] {
		seen[n] = true
		n = next(n)
	}
	return n == 1
}

func next(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		sum += d * d
	}
	return sum
}

// 202. Happy Number
// 202. 快乐数
// 思路：快慢指针
// time O(logn) space O(1)
func isHappy2(n int) bool {
	slow, fast := n, next(n)
	for fast != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}
