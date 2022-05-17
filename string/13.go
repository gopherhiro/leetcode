package string

// 13. Roman to Integer
// 13. 罗马数字转整数
// 思路：根据罗马数字与整数规律即可。
// time O(N) space O(1)
func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ans, n := 0, len(s)
	for i := 0; i < n; i++ {
		curr := romanMap[s[i]]
		// 若数字比它右侧的数字小，则将该数字的符号取反。
		if i < n-1 && curr < romanMap[s[i+1]] {
			curr = -curr
		}
		ans += curr
	}
	return ans
}
