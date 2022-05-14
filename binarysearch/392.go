package binarysearch

// 392. 判断子序列
// 392. Is Subsequence
// 思路：双指针 time (N + M),space O(1)
// 初始化两个指针 i 和 j，分别指向 s 和 t 的初始位置；
// 每次贪心地匹配，匹配成功则 i 和 j 同时右移，匹配 s 的下一个位置；
// 匹配失败，则 j 右移，i 不变，尝试用 t 的下一个字符匹配 s；
// 最终如果 i 移动到 s 的末尾，就说明 s 是 t 的子序列。
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}
