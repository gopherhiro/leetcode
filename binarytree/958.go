package binarytree

// 958. 二叉树的完全性检验
// 958. Check Completeness of a Binary Tree
// 策略：BFS
// 思路：对于一颗完全二叉树，层序遍历的过程中，遇到第一个空节点之后不应该再出现非空节点。
// Use BFS to do a level order traversal,
// add children to the bfs queue,
// until we met the first empty node.
// For a complete binary tree,
// there should not be any node after we met an empty one.
// time O() space O()
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	reachEnd := false
	queue := make([]*TreeNode, 0)
	// 根节点入队
	queue = append(queue, root)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// 遇到第一个空节点之后，再次出现了非空节点
		// 则不是完全二叉树
		if reachEnd && curr != nil {
			return false
		}

		// 标记遇到第一个空节点
		if curr == nil {
			reachEnd = true
			continue
		}
		queue = append(queue, curr.Left)
		queue = append(queue, curr.Right)
	}

	return true
}
