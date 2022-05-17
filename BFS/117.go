package BFS

// 117. Populating Next Right Pointers in Each Node II
// 117. 填充每个节点的下一个右侧节点指针 II
// 思路：BFS-层序遍历
// time O(N) space O(M)
func connectII(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var prev *TreeNode
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			cur := queue[0]
			queue = queue[1:]

			if prev != nil {
				prev.Next = cur
			}
			prev = cur

			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
	}

	return root
}
