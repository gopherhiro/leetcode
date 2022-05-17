package tree

// 105. Construct Binary TreeNode from Preorder and Inorder Traversal
// 思路：构造二叉树，第一件事一定是找根节点，然后想办法构造左右子树。
// 通过前序或者后序遍历结果找到根节点，然后根据中序遍历结果确定左右子树。
// time O(N^2)
// space O(N)
func buildTreeFromPreIn(preorder []int, inorder []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// root 节点对应的值就是前序遍历数组的第一个元素
	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	// rootValue 在中序遍历数组中的索引
	inorderIndex := findInorderIndex(inorder, rootValue)

	// 递归调用构建左右子树 : 这里的下标需要计算
	root.Left = buildTreeFromPreIn(preorder[1:inorderIndex+1], inorder[:inorderIndex])
	root.Right = buildTreeFromPreIn(preorder[inorderIndex+1:], inorder[inorderIndex+1:])
	return root
}

func findInorderIndex(inorder []int, value int) int {
	for i, v := range inorder {
		if v == value {
			return i
		}
	}
	return 0
}
