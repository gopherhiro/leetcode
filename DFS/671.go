package DFS

// 671. Second Minimum Node In a Binary Tree
// 671. 二叉树中第二小的节点
// 思路：DFS-有条件的前序遍历
// time O(N) space O(N)
func findSecondMinimumValue(root *TreeNode) int {
	// base case
	// 00
	if root.Left == nil && root.Right == nil {
		return -1
	}
	// 根据题目要求，01、10 case 不会存在。
	// 但为了运行，需要做一下处理。
	// 01
	if root.Left == nil && root.Right != nil {
		return root.Right.Val
	}
	// 10
	if root.Left != nil && root.Right == nil {
		return root.Left.Val
	}

	// 11
	left, right := root.Left.Val, root.Right.Val

	// 如果左右子节点都等于 root.val，则去左右子树递归寻找第二小的值
	if root.Val == left {
		left = findSecondMinimumValue(root.Left)
	}
	if root.Val == right {
		right = findSecondMinimumValue(root.Right)
	}

	if left == -1 {
		return right
	}
	if right == -1 {
		return left
	}

	// 如果左右子树都找到一个第二小的值，更小的那个是整棵树的第二小元素
	return min(left, right)
}
