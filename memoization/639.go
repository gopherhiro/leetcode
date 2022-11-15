package memoization

// 639. Decode Ways II
// 639. 解码方法 II
// 思路：memoization
// time O(n) space O(n)
func numDecodingsII(s string) int {
	const M int = 1e9 + 7
	n := len(s)
	m := make(map[int]int, n+1)
	var ways func(i int) int
	ways = func(i int) int {
		// base case
		if i < 0 {
			return 1
		}
		// query from memory
		if c, ok := m[i]; ok {
			return c
		}

		// special handling : *
		// s[i] == '*'
		if s[i] == '*' {
			count := 9 * ways(i-1) % M
			if i > 0 && s[i-1] == '1' {
				count = (count + 9*ways(i-2)) % M
			}
			if i > 0 && s[i-1] == '2' {
				count = (count + 6*ways(i-2)) % M
			}
			if i > 0 && s[i-1] == '*' {
				count = (count + 15*ways(i-2)) % M
			}
			m[i] = count
			return count
		}

		// s[i] != '*'
		count := 0
		if s[i] == '0' {
			count = 0
		} else {
			count = ways(i - 1)
		}

		if i > 0 && s[i-1] == '1' {
			count = (count + ways(i-2)) % M
		}

		if i > 0 && s[i-1] == '2' && s[i] <= '6' {
			count = (count + ways(i-2)) % M
		}

		if i > 0 && s[i-1] == '*' {
			if s[i] <= '6' {
				// s[i-1] can be 1,2
				count = (count + 2*ways(i-2)) % M
			} else {
				// s[i-1] only be 1
				count = (count + ways(i-2)) % M
			}
		}
		m[i] = count
		return count
	}
	return ways(n - 1)
}
