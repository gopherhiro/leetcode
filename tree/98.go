package tree

import "math"

// 98. Validate Binary Search Tree
// DFS-post order traversal：Recursion, time O(N), space O(N)
// 思路：递归验证根节点是否在相应的范围内
// root.val => (min, max)
// root.left.val => (min, root.val)
// root.right.val => (root.val, max)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var traverse func(root *TreeNode, min, max int) bool
	traverse = func(root *TreeNode, min, max int) bool {
		// base case
		if root == nil {
			return true
		}
		// root的值，必须满足：min < root.Val < max
		if root.Val <= min || root.Val >= max {
			return false
		}

		// 左子树的最大值是 root.val
		leftIsBST := traverse(root.Left, min, root.Val)

		// 右子树的最小值是 root.val
		rightIsBST := traverse(root.Right, root.Val, max)

		// left is BST and right is BST, so the tree is BST
		return leftIsBST && rightIsBST
	}
	return traverse(root, math.MinInt64, math.MaxInt64)
}
