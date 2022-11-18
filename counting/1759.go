package counting

// 1759. Count Number of Homogenous Substrings
// 1759. 统计同构子字符串的数目
// 思路：one pass
// time O(n) space O(1)
func countHomogenous(s string) int {
	const M int = 1e9 + 7
	res, count := 0, 0
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			count += 1
		} else {
			count = 1
		}
		res = (res + count) % M
	}
	return res
}
