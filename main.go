package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, -3, -5}
	r := arraySign(nums)
	fmt.Println(r)
}

// 1822. Sign of the Product of an Array
// 1822. 数组元素积的符号
// 思路：负负得正
// time O(n) space O(1)
func arraySign(nums []int) int {
	ans := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		if num < 0 {
			ans = -ans
		}
	}
	return ans
}
