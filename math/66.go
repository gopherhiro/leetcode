package math

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
