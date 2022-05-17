package tree

import "math"

// 654. Maximum Binary TreeNode
// 基本思路：先要找到根节点，然后让 build 函数递归生成左右子树即可。
func constructMaxBinaryTree(nums []int) *TreeNode {
	root := build(nums, 0, len(nums)-1)
	return root
}

func build(nums []int, begin, end int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if begin > end {
		return nil
	}

	// 找到数组中最大值及其索引
	maxValue, index := math.MinInt8, -1
	for i := begin; i <= end; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			index = i
		}
	}

	// 构建根结点
	root := &TreeNode{Val: maxValue}

	// 递归调用构建左右子树
	root.Left = build(nums, begin, index-1)
	root.Right = build(nums, index+1, end)

	return root
}

// 654. Maximum Binary TreeNode
// 基本思路：先要找到根节点，然后让 build 函数递归生成左右子树即可。
// time O(N^2)
// space O(N)
func constructMaximumBinaryTree(nums []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	// 找到数组中最大值的索引，构建二叉树节点
	index := findMaxValueIndex(nums)
	root := &TreeNode{Val: nums[index]}

	// 递归调用构建左右子树
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func findMaxValueIndex(nums []int) int {
	maxValueIndex := 0
	for i, v := range nums {
		if v > nums[maxValueIndex] {
			maxValueIndex = i
		}
	}
	return maxValueIndex
}
