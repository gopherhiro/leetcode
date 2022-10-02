package main

import "fmt"

func main() {
	r := countBits2(3)
	fmt.Println(r)
}

func countBits2(n int) []int {
	ans := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 二进制表示中
		// 奇数中 1 的个数，比前面那个偶数多一个 1
		// 偶数中 1 的个数，和除以 2 之后的那个数一样多
		if i%2 == 1 {
			ans[i] = ans[i-1] + 1
		} else {
			ans[i] = ans[i/2]
		}
	}
	return ans
}

func countBits(n int) []int {
	ans := make([]int, 0, n+1)
	for i := 0; i <= n; i++ {
		res := countBit(i)
		ans = append(ans, res)
	}
	return ans
}

func countBit(num int) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
