package tree

// 103. Binary TreeNode Zigzag Level Order Traversal
// 思路：与二叉树层序遍历一样，只要再加一个bool变量控制遍历方向即可。
// time O(N), space O(N)
func zigzagLevelOrder(t *TreeNode) [][]int {
	if t == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// 方向标识：向右
	toRight := true

	// 根节点入队
	queue = append(queue, t)

	// 第一层循环控制：从上向下层层遍历
	for len(queue) > 0 {
		// 记录这一层的节点值
		row := make([]int, 0)
		rowSize := len(queue)

		// 第二层循环控制每一层：从左向右遍历
		for i := 0; i < rowSize; i++ {
			// 从队列中弹出一个节点处理
			cur := queue[0]
			queue = queue[1:]

			// 实现之字遍历
			if toRight {
				row = append(row, cur.Val)
			} else {
				row = append([]int{cur.Val}, row...)
			}

			// 其有左右子节点，则左右子节点入队
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 切换方向
		toRight = !toRight
		res = append(res, row)
	}

	return res
}
