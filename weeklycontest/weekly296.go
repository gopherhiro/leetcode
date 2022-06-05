package main

import "fmt"

func main() {
	nums := []int{1, 2}
	operations := [][]int{
		{1, 3},
		{2, 1},
		{3, 2},
	}
	r := arrayChange(nums, operations)
	fmt.Println(r)
}

// 6092. 替换数组中的元素
// 思路：哈希表
// time O(m+n) space O(n)
func arrayChange(nums []int, operations [][]int) []int {
	ht := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ht[nums[i]] = i
	}
	m := len(operations)
	for i := 0; i < m; i++ {
		opi0 := operations[i][0]
		opi1 := operations[i][1]
		idx := ht[opi0]
		nums[idx] = opi1
		ht[opi1] = idx
	}
	return nums
}

// 6090. 极大极小游戏
// 思路：模拟
// time O(N*logN) space O(logN)
func minMaxGame(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	for len(nums) > 1 {
		mid := len(nums) / 2
		newNums := make([]int, mid)
		for i := 0; i < mid; i++ {
			if i%2 == 0 {
				newNums[i] = min(nums[2*i], nums[2*i+1])
			} else {
				newNums[i] = max(nums[2*i], nums[2*i+1])
			}
		}
		nums = newNums
	}
	return nums[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
