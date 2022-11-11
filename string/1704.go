package string

// 1704. Determine if String Halves Are Alike
// 1704. 判断字符串的两半是否相似
// 思路：hash table
// time O(n) space O(1)
func halvesAreAlike(s string) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	m := make(map[byte]bool, 0)
	for _, v := range vowels {
		m[v] = true
	}
	mid := len(s) / 2
	return count(s[:mid], m) == count(s[mid:], m)
}

func count(s string, m map[byte]bool) int {
	c := 0
	for _, v := range s {
		if m[byte(v)] {
			c++
		}
	}
	return c
}

// 1704. Determine if String Halves Are Alike
// 1704. 判断字符串的两半是否相似
// 思路：hash table
// time O(n) space O(1)
func halvesAreAlikeN(s string) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	m := make(map[byte]bool, 0)
	for _, v := range vowels {
		m[v] = true
	}
	n := len(s)
	mid := n / 2
	count := 0
	for i := 0; i < n; i++ {
		if !m[s[i]] {
			continue
		}
		// front part count increase by 1
		// back part count decrease by 1
		// if count is 0, halves are like
		if i < mid {
			count += 1
		} else {
			count -= 1
		}
	}
	return count == 0
}

func halvesAreAlikeNice(s string) bool {
	mid, cnt := len(s)/2, 0
	for i, c := range s {
		switch c {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			if i < mid {
				cnt++
			} else {
				cnt--
			}
		}
	}
	return cnt == 0
}
