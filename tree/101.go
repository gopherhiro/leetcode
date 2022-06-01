package tree

// 101. 对称二叉树
// 策略：DFS
// time O(N) space O(N)
func isSymmetric(root *TreeNode) bool {
	var helper func(p, q *TreeNode) bool
	helper = func(p, q *TreeNode) bool {
		// base case
		// 前序遍历位置
		if p == nil && q == nil {
			return true
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		return helper(p.Left, q.Right) && helper(p.Right, q.Left)
	}
	return helper(root, root)
}

// 101. 对称二叉树
// 策略：迭代，使用队列辅助
// time O(N) space O(N)
func isSymmetricI(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root, root)

	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil && q == nil {
			continue
		}
		if p == nil || q == nil {
			return false
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Right)
		queue = append(queue, p.Right, q.Left)
	}
	return true
}
