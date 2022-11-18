package counting

// 1446. Consecutive Characters
// 1446. 连续字符
// 思路：一次遍历
// time O(n) space O(1)
func maxPower(s string) int {
	if len(s) == 0 {
		return 0
	}
	power, counter := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			counter++
		} else {
			counter = 1
		}
		power = max(power, counter)
	}
	power = max(power, counter)
	return power
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
