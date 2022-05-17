package tree

// 226. Invert Binary TreeNode
// time O(N): N 为二叉树节点的数目
// space O(N): 使用的空间由递归栈的深度决定
// 思路：遍历二叉树的所有节点，并且对每个节点实施「交换子节点」的操作。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	/*** 前序位置 ***/
	// 交换 root 节点的左右子节点
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	// 左右子节点继续翻转它们的子节点
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 方案2：分解问题解法。
// 思路：左子树的节点翻转，右子树的节点翻转，然后再交换左右子树。
func invertTreeN(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTreeN(root.Left)
	right := invertTreeN(root.Right)
	root.Left = right
	root.Right = left
	return root
}
