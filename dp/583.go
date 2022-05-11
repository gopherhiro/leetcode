package dp

// 583. Delete Operation for Two Strings
// 583. 两个字符串的删除操作
// 思路：两个字符串的长度分别减去[最长公共子序列]长度，再求和即可。
// time O(m*n) space O(m*n)
func minDeleteStep(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}
	lcs := longestCommonSubsequence(word1, word2)
	// 步数即是：两个单词变为其最长公共子序列过程，所需要删除的字符个数
	return m - lcs + n - lcs
}
