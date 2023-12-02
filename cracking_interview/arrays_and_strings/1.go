package arrays_and_strings

// 题目：判定字符是否唯一
// 思路：哈希表
// 复杂度：time O(n), space O(n)
func isUnique(s string) bool {
	m := make(map[byte]bool, len(s))
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if m[ch] {
			return false
		}
		m[ch] = true
	}
	return true
}

func isUniqueC(s string) bool {
	charSet := make(map[rune]bool)
	for _, char := range s {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

// 题目：判定字符是否唯一
// 思路：位运算
// 复杂度：time O(n), space O(1)
//我们可以把它看成000...00(26个0)，这26个bit对应着26个字符，对于一个字符c，检查对应下标的bit值即可判断是否重复。
//那么难点在于如何检查？这里我们可以通过位运算来完成。
//首先计算出字符char离'a'这个字符的距离，即我们要位移的距离，用move_bit表示，
//那么使用左移运算符1 << move_bit则可以得到对应下标为1，其余下标为0的数。
//如字符char = 'c'，则得到的数为000...00100，将这个数跟 checker 做与运算，
//由于这个数只有一个位为1，其他位为0，那么与运算的结果中，其他位肯定是0，
//而对应的下标位是否0则取决于之前这个字符有没有出现过。
//若出现过则被标记为1，那么与运算的结果就不为0；
//若之前没有出现过，则对应位的与运算的结果也是0，那么整个结果也为0。
//对于没有出现过的字符，我们用或运算 checker | (1 << move_bit)将对应下标位的值置为1。
func isUniqueB(s string) bool {
	checker := 0
	for _, ch := range s {
		moveBit := ch - 'a'
		if checker&(1<<moveBit) > 0 {
			return false
		}
		checker |= 1 << moveBit
	}
	return true
}
