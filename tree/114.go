package tree

// 114. Flatten Binary TreeNode to Linked List
// time O(N^2) 遍历的二叉树节点个数 O(N)，最后右子树的拼接 O(N)。
// space O(N) 取决于函数递归调用栈。
func flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := flatten(root.Left)
	right := flatten(root.Right)

	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表
	// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

	return root
}
