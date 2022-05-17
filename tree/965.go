package tree

// 965. Univalued Binary TreeNode
// Approach 1: Depth-First Search
// time O(N), space O(N)
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	isUnival := true

	// 记下第一个节点的值
	// 遍历二叉树，看看是否存在其他值。
	first := root.Val
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}
		if first != root.Val {
			isUnival = false
			return
		}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)

	return isUnival
}

// Approach 2: Recursion
// 分解问题：左右子树都是单值的，则该树是单值的。
// time O(N), space O(Height)
func isUnivalTreeRecur(root *TreeNode) bool {
	// 左子树是单值的
	leftIsUnival := root.Left == nil || root.Val == root.Left.Val && isUnivalTreeRecur(root.Left)
	// 右子树是单值的
	rightIsUnival := root.Right == nil || root.Val == root.Right.Val && isUnivalTreeRecur(root.Right)
	// 从而该二叉树是单值的。
	return leftIsUnival && rightIsUnival
}
