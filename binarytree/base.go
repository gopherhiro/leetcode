package binarytree

import "fmt"

// preorder traverse tree
func preorderTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	// preorder position
	fmt.Printf("%d ", t.Val)
	preorderTraverse(t.Left)
	preorderTraverse(t.Right)
}

// inorder traverse tree
func inorderTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	inorderTraverse(t.Left)
	fmt.Printf("%d ", t.Val)
	inorderTraverse(t.Right)
}

// post order traverse tree
func posOrderTraverse(t *TreeNode) {
	if t == nil {
		return
	}
	posOrderTraverse(t.Left)
	posOrderTraverse(t.Right)
	fmt.Printf("%d ", t.Val)
}

func layerOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	// 根节点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		// 已遍历节点出队
		cur := queue[0]
		queue = queue[1:]
		fmt.Printf("%d ", cur.Val)

		// 其左右子节点入队
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}

		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}
	}
}

// 输入一棵二叉树的根节点，层序遍历这棵二叉树
func levelOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	// 从上到下遍历二叉树的每一层
	for len(queue) > 0 {
		// 从左到右遍历每一层的每个节点
		rowSize := len(queue)
		for i := 0; i < rowSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			fmt.Printf("%d ", cur.Val)

			// 将下一层节点放入队列
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}
}
