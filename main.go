package main

import (
	"fmt"
	"math/bits"
)

func main() {
	r := minimizeXor(3, 5)
	fmt.Println(r)
}

// 2429. 最小 XOR
// 思路：位运算
// time O(logn) space O(1)
func minimizeXor(num1 int, num2 int) int {
	c1 := bits.OnesCount(uint(num1))
	c2 := bits.OnesCount(uint(num2))
	if c1 == c2 {
		return num1
	}
	for ; c1 < c2; c2-- {
		num1 = num1 | (num1 + 1) // 最低位上的 0 变为 1
	}
	for ; c1 > c2; c2++ {
		num1 = num1 & (num1 - 1) // 最低位上的 1 变为 0
	}
	return num1
}
