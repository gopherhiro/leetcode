package counting

// 1513. Number of Substrings With Only 1s
// 1513. 仅含 1 的子串数
// 思路：一次遍历
// time O(n) space O(1)
func numSub(s string) int {
	const M int = 1e9 + 7
	res, count := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}
		if i > 0 && s[i] == s[i-1] {
			count += 1
		} else {
			count = 1
		}
		res = (res + count) % M
	}
	return res
}
