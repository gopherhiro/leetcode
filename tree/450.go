package tree

// 450. Delete Node in a BST
// Recursion, time O(H), space O(H)
// 删除时，注意区分情况：case 1，case 2，case 3，对应不同的删除操作。
func deleteNodeBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到了要删除的节点
	if val == root.Val {
		// case 1，要删除的节点是叶子节点（case 1可以删除，因为case 2情况已包含该情况）
		// 为了可读性，这里可以保留。
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// case 2，要删除的节点只有一个子节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// case 3，要删除的节点有两个子节点
		// root.left != nil && root.right != nil
		// 找到右子树的最小节点
		minNode := getMin(root.Right)
		// 删除右子树最小的节点
		root.Right = deleteNodeBST(root.Right, minNode.Val)
		// 用右子树最小的节点替换 root 节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	}

	// 到左子树中删除
	if val < root.Val {
		root.Left = deleteNodeBST(root.Left, val)
	}

	// 到右子树中删除
	if val > root.Val {
		root.Right = deleteNodeBST(root.Right, val)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// BST 最左边的就是最小的。
	for root.Left != nil {
		root = root.Left
	}
	return root
}
