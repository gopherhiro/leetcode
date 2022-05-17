package tree

// 889. Construct Binary TreeNode from Preorder and Postorder Traversal
// 思路：构造二叉树，第一件事一定是找根节点，然后想办法构造左右子树。
// 通过前序或者后序遍历结果找到根节点，然后根据中序遍历结果确定左右子树。
// time O(N^2)
// space O(N)
func buildTreeFromPrePost(preorder, postorder []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// root 节点对应的值就是前序遍历数组的第一个元素
	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	// 把前序遍历结果的第二个元素作为左子树的根节点的值。
	leftRootValue := preorder[1]
	index := rootLeftIndex(postorder, leftRootValue)

	// 在后序遍历结果中寻找左子树根节点的值，从而确定了左子树的索引边界，进而确定右子树的索引边界
	// 从而递归构造左右子树即可。
	root.Left = buildTreeFromPrePost(preorder[1:index+1], postorder[:index+1])
	root.Right = buildTreeFromPrePost(preorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func rootLeftIndex(postorder []int, value int) int {
	for i, v := range postorder {
		if v == value {
			// 这里返回的索引需要+1
			return i + 1
		}
	}
	return 0
}
