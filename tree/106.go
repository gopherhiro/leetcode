package tree

// 106. Construct Binary TreeNode from Inorder and Postorder Traversal
// 思路：构造二叉树，第一件事一定是找根节点，然后想办法构造左右子树。
// 通过前序或者后序遍历结果找到根节点，然后根据中序遍历结果确定左右子树。
// time O(N^2)
// space O(N)
func buildTreeFromInPost(inorder []int, postorder []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	// root 节点对应的值就是后序遍历数组的最后一个元素
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootValue}

	// rootValue 在中序遍历数组中的索引
	inorderIndex := findInorderIndex(inorder, rootValue)

	root.Left = buildTreeFromInPost(inorder[:inorderIndex], postorder[:inorderIndex])
	root.Right = buildTreeFromInPost(inorder[inorderIndex+1:], postorder[inorderIndex:len(postorder)-1])

	return root
}
