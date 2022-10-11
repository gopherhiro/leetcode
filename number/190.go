package math

// 190. Reverse Bits
// 190. 颠倒二进制位
// 思路：循环（位运算）
// time O(1) space O(1)
func reverseBits(num uint32) uint32 {
	// 模拟运算：把 num 的 bits 颠倒过来。
	var ans uint32
	for i := 0; i < 32; i++ {
		ans = ans << 1 // ans 左移
		ans += num & 1 // num 的最低位是否为1，是 ans++，否 ans 不变。
		num = num >> 1 // num 右移
	}
	return ans
}
