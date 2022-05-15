package twopointer

// 345. Reverse Vowels of a String
// 345. 反转字符串中的元音字母
// 思路：双指针
// time O(N) space O(1)
func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}

	t := []byte(s)

	left, right := 0, len(t)-1
	for left < right {
		// 找到左边的元音字母
		for left < right && !isVowel(t[left]) {
			left++
		}
		// 找到右边的元音字母
		for left < right && !isVowel(t[right]) {
			right--
		}
		// left, right 都找到，则进行交换
		// 交换之后继续遍历
		t[left], t[right] = t[right], t[left]
		left++
		right--
	}

	return string(t)
}

func isVowel(ch byte) bool {
	return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
		ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U'
}
