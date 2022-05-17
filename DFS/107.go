package DFS

// 107. Binary Tree Level Order Traversal II
// 107. 二叉树的层序遍历 II
// 思路：层序遍历框架
// time O(N), space O(M)
func levelOrderBottom(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 把遍历得到的最新一层放到头部，即是自底向上的遍历。
		ans = append([][]int{level}, ans...)
	}
	return
}

// 107. Binary Tree Level Order Traversal II
// 107. 二叉树的层序遍历 II
// 思路：层序遍历框架
// time O(N) space O(M)
func levelOrderBottomN(root *TreeNode) (ans [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		level := make([]int, 0)
		levelSize := len(queue)

		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			level = append(level, cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		ans = append(ans, level)
	}

	// 颠倒一下元素顺序
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-1-i] = ans[len(ans)-1-i], ans[i]
	}
	return
}
