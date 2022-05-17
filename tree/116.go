package tree

// 116. 填充每个节点的下一个右侧节点指针
// Populating Next Right Pointers in Each Node
// time O(N) 遍历二叉树的节点
// space O(N) 函数递归调用栈
func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(left *TreeNode, right *TreeNode) {
	if left == nil || right == nil {
		return
	}
	/**** 前序遍历位置 ****/
	// 将传入的两个节点连接
	left.Next = right

	// 连接相同父节点的两个子节点
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(right.Left, right.Right)

	// 连接跨越父节点的两个子节点
	connectTwoNode(left.Right, right.Left)
}

// 116. 填充每个节点的下一个右侧节点指针
// Populating Next Right Pointers in Each Node
func addLinkToTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	addLink(root)
	return root
}
func addLink(root *TreeNode) {
	if root.Left == nil || root.Right == nil {
		return
	}

	/**** 前序遍历位置 ****/
	// 连接相同父节点的两个子节点
	root.Left.Next = root.Right

	// 连接跨越父节点的两个子节点
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	addLink(root.Left)
	addLink(root.Right)
}
