package tree

// 111. Minimum Depth of Binary TreeNode
// 思路：BFS（Breadth-First Search）
// 从上到下，从左至右，第一个达到叶子节点的层，即是树的最小深度。
// BFS与DFS的区别，BFS第一次搜索到的结果是最优的。
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断是否到达叶子结点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			// 层序遍历的框架
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 更新深度
		depth++
	}
	return depth
}
