package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9},
	}
	r := maxSum(grid)
	fmt.Println(r)
}

// 6193. 沙漏的最大总和
// 思路：模拟统计
// time O(mn) space O(1)
func maxSum(grid [][]int) int {
	ans := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			up := grid[i-1][j-1] + grid[i-1][j] + grid[i-1][j+1]
			down := grid[i+1][j-1] + grid[i+1][j] + grid[i+1][j+1]
			sum := up + grid[i][j] + down
			ans = max(ans, sum)
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
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
