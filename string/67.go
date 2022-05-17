package string

// 67. Add Binary
// 67. 二进制求和
// 思路：模拟加法
// time (max(m,n)) space O(max(m,n))
func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	ans := make([]byte, 0)
	carry := 0 // 进位
	i, j := len(a)-1, len(b)-1
	for i >= 0 || j >= 0 {
		va, vb := 0, 0
		if i >= 0 {
			va = int(a[i] - '0')
			i--
		}
		if j >= 0 {
			vb = int(b[j] - '0')
			j--
		}
		sum := va + vb + carry // 求和
		carry = sum / 2        // 进位
		curr := sum % 2        // 当前值
		res := byte(curr + '0')
		ans = append([]byte{res}, ans...)
	}
	// 如果最后有进位，则需要再加一位
	if carry > 0 {
		last := byte(carry + '0')
		ans = append([]byte{last}, ans...)
	}
	return string(ans)
}
