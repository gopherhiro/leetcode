package main

import (
	"fmt"
	"math"
)

func main() {
	r := minimizeXor(25, 72)
	fmt.Println(r)
}

// 6194. 最小 XOR
// 思路：位运算
// time O(nlogn) space O(1)
// 提交测试：时间超限
func minimizeXor(num1 int, num2 int) int {
	ans := 0
	minor := math.MaxInt
	count := countBit(num2)
	fmt.Println("count:", count)
	num := getNum(count)
	fmt.Println("num:", num)
	for i := 1; i <= num; i++ {
		if countBit(i) != count {
			continue
		}
		xor := i ^ num1
		if xor < minor {
			minor = xor
			ans = i
		}
	}
	return ans
}

func getNum(n int) int {
	var ans float64
	for i := 0; i < n; i++ {
		ans += math.Pow(2, float64(i))
	}
	return int(ans)
}

// time O(logn) space O(1)
func countBit(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

// 6192. 公因子的数目
// 思路：模拟统计
// time O(n) space O(1)
func commonFactors(a int, b int) int {
	count := 0
	m := min(a, b)
	for i := 1; i <= m; i++ {
		if a%i == 0 && b%i == 0 {
			count++
		}
	}
	return count
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
