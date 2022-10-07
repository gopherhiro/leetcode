package main

import "fmt"

func main() {
	A := []int{5, 10, 16, -6, 15, 11, 3}
	B := []int{16, 22, 23, 23, 25, 3, -16}
	r := temperatureTrend(A, B)
	fmt.Println(r)
}

// 1. 气温变化趋势
// 对于第 i ~ (i+1) 天的气温变化趋势，将根据以下规则判断：
// 若第 i+1 天的气温 高于 第 i 天，为 上升 趋势
// 若第 i+1 天的气温 等于 第 i 天，为 平稳 趋势
// 若第 i+1 天的气温 低于 第 i 天，为 下降 趋势
// 已知 temperatureA[i] 和 temperatureB[i] 分别表示第 i 天两地区的气温。
// 组委会希望找到一段天数尽可能多，且两地气温变化趋势相同的时间举办嘉年华活动。
// 请分析并返回两地气温变化趋势相同的最大连续天数。

func temperatureTrend(A []int, B []int) int {
	ra, rb := make([]int, 0), make([]int, 0)
	length := len(A)
	for i := 1; i < length; i++ {
		a := compare(A[i-1], A[i])
		ra = append(ra, a)
		b := compare(B[i-1], B[i])
		rb = append(rb, b)
	}

	res, seqLen := 0, 0
	for i := 0; i < length-1; i++ {
		if ra[i] == rb[i] {
			seqLen++
			continue
		} else {
			res = max(res, seqLen)
			seqLen = 0
		}
	}
	res = max(res, seqLen)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func compare(before, after int) int {
	if after == before {
		return 0
	}
	if after > before {
		return 1
	}
	return -1
}

// 最小展台数量
// 思路：哈希表
// time O(mn) space O(1)
// https://leetcode.cn/contest/season/2022-fall/problems/600YaG/
func minNumBooths(demand []string) int {
	ht := make(map[rune]int, 0)
	for _, str := range demand {
		m := make(map[rune]int, 0)
		for _, ch := range str {
			m[ch]++
		}
		for char, count := range m {
			v, ok := ht[char]
			// 不存在
			if !ok {
				ht[char] = count
				continue
			}
			// 存在
			if ok && count > v {
				ht[char] = count
			}
		}
	}
	sum := 0
	for _, count := range ht {
		sum += count
	}
	return sum
}

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
	Next  *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 装饰树
// 思路：BFS
// time O(n) space O(n)
// https://leetcode.cn/contest/season/2022-fall/problems/KnLfVT/
func expandBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
				new := &TreeNode{Val: -1}
				tmp := cur.Left
				cur.Left = new
				new.Left = tmp
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
				new := &TreeNode{Val: -1}
				tmp := cur.Right
				cur.Right = new
				new.Right = tmp
			}
		}
	}
	return root
}

// 美观的花束
// 思路：回溯
func beautifulBouquet(flowers []int, target int) [][]int {
	res := make([][]int, 0)

	// 记录已选择路径
	track := make([]int, 0)

	var backtrack func(nums []int, start int)
	backtrack = func(nums []int, start int) {
		// 前序位置：收集结果。
		tmp := make([]int, len(track))
		copy(tmp, track)

		// 判断条件
		if !isBeautiful(tmp, target) {
			return
		}

		res = append(res, tmp)

		// 使用 start 参数控制递归，进行剪枝。
		for i := start; i < len(nums); i++ {

			// 剪枝。值相同的相邻树枝，只遍历第一条。
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// 做选择
			track = append(track, nums[i])

			// 进入下一层抉择树
			backtrack(nums, i+1)

			// 撤销选择
			track = track[:len(track)-1]
		}

	}
	backtrack(flowers, 0)
	return res
}

func isBeautiful(flowers []int, count int) bool {
	ht := make(map[int]int, 0)
	for _, v := range flowers {
		ht[v]++
		if ht[v] > count {
			return false
		}
	}
	return true
}
