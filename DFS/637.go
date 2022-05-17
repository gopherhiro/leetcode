package DFS

// 637. Average of Levels in Binary Tree
// 637. 二叉树的层平均值
// 思路：BFS-层序遍历
// time O(N) space O(N)
func averageOfLevels(root *TreeNode) (ans []float64) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		sum := float64(0)
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			sum += float64(cur.Val)

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		r := sum / float64(levelSize)
		ans = append(ans, r)

	}
	return
}
