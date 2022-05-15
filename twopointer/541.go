package twopointer

import "leetcode/pkg"

// 541. Reverse String II
// 541. 反转字符串 II
// 思路：双指针
// time (N) space O(N)
func reverseStr(s string, k int) string {
	if len(s) == 0 || k == 0 {
		return s
	}

	t := []byte(s)

	for i := 0; i < len(t); i += 2 * k {
		// 计算交换范围
		// right = i+k-1，表示右边界范围
		// right 与 len(t) - 1 取 min，
		// 可保证不足 k 个时，将剩余字符全部反转。
		left, right := i, pkg.Min(i+k-1, len(t)-1)
		for left < right {
			t[left], t[right] = t[right], t[left]
			left++
			right--
		}
	}

	return string(t)
}

// Reverse Kth String
// 反转前k个字符
// 思路：左右指针，交换k-1个元素即可。
// time O(N) space (1)
func reverseKthStr(s string, k int) string {
	if len(s) == 0 || k == 0 {
		return s
	}
	ns := []byte(s)
	left, right := 0, k-1
	for left < right {
		ns[left], ns[right] = ns[right], ns[left]
		left++
		right--
	}
	return string(ns)
}
