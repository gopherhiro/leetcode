package string

import "bytes"

// 481. Magical String
// 481. 神奇字符串
// 思路：模拟(关键是弄懂如何生成"神奇字符串")
// time O(n) space O(n)
func magicalString(n int) int {
	s := make([]byte, 0, n+1)
	s = append(s, 1, 2, 2)
	c := byte(2)
	for i := 2; len(s) < n; i++ {
		// x ^ y = z
		// z ^ x = y
		// z ^ y = x
		// 1 ^ 2 = 3, so that
		// 1 ^ 3 = 2
		// 2 ^ 3 = 1
		c = c ^ 3 // 1 => 2, 2 => 1
		count := int(s[i])
		b := bytes.Repeat([]byte{c}, count)
		s = append(s, b...)
	}
	return bytes.Count(s[:n], []byte{1})
}
