package tree

import "math"

// 222. Count Complete TreeNode Nodes
// 思路：遍历完全二叉树，统计节点个数即可。
// time O(N), space O(N)
var treeNodesNumber = 0

func countTreeNodes(root *TreeNode) int {
	countNodesHelper(root)
	return treeNodesNumber
}

func countNodesHelper(root *TreeNode) {
	if root == nil {
		return
	}
	treeNodesNumber++
	countNodesHelper(root.Left)
	countNodesHelper(root.Right)
}

// 222. Count Complete TreeNode Nodes
// 思路：满二叉树节点计算公式+递归。
// time (logN^2), space O(N)
func countTreeNodesOptimal(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 记录左、右子树的高度
	left, right := root, root
	leftHigh, rightHigh := 0, 0
	for left != nil {
		left = left.Left
		leftHigh++
	}
	for right != nil {
		right = right.Right
		rightHigh++
	}
	// 如果左右子树的高度相同，则是一棵满二叉树
	// 满二叉树的节点个数，可以直接通过树高算出来。
	if leftHigh == rightHigh {
		nodesNumber := math.Pow(float64(2), float64(leftHigh)) - 1
		return int(nodesNumber)
	}
	// 如果左右高度不同，则按照普通二叉树的逻辑计算
	return countTreeNodesOptimal(root.Left) + countTreeNodesOptimal(root.Right) + 1
}
