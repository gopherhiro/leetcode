package BFS

// 404. Sum of Left Leaves
// 404. 左叶子之和
// 思路：BFS-层序遍历。遍历二叉树，累计左子节点之和即可。
// 左子节点：一个节点是左节点，并且是叶子节点。
// time O(N) space O(N)
func sumOfLeftLeavesBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 判断是否是叶子节点
	isLeaf := func(root *TreeNode) bool {
		return root.Left == nil && root.Right == nil
	}

	sum := 0

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		if cur.Left != nil && isLeaf(cur.Left) {
			sum += cur.Left.Val
		}

		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

	}
	return sum
}
