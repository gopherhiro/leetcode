package tree

// 230. Kth Smallest Element in a BST
// DFS-Recursive Inorder Traversal，time O(N), space O(N)
// 思路：BST，升序排序，然后找第 k 个元素。
func kthSmallest(root *TreeNode, k int) int {
	// 记录当前元素的排名
	var rank int
	// 记录结果
	var answer int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		/* 中序位置 */
		// 找到第k小的元素
		rank++
		if rank == k {
			answer = root.Val
			return
		}
		traverse(root.Right)
	}
	traverse(root)
	return answer
}

// Approach 2: Iterative Inorder Traversal
// The above recursion could be converted into iteration, with the help of stack.
// This way one could speed up the solution
// because there is no need to build the entire inorder traversal,
// and one could stop after the kth element.
// time (H+k), space (H)
func kthSmallestIter(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 中序位置
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		// 找到第k小的元素
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
