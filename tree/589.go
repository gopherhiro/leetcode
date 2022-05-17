package tree

// 589. N-ary TreeNode Preorder Traversal
// 思路：使用N叉树的前序遍历框架
// DFS(depth first search), time O(N), space O(N)
func preorderOfNTree(root *Node) []int {
	answer := make([]int, 0)
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		// 前序遍历位置
		answer = append(answer, root.Val)
		for _, child := range root.Children {
			traverse(child)
		}
		// 后序遍历位置
	}
	traverse(root)
	return answer
}

// 迭代，利用栈。
// time O(N) space O(N)
func preorderOfNTreeIter(root *Node) (answer []int) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		answer = append(answer, node.Val)

		// 首先把根节点入栈，因为根节点是前序遍历中的第一个节点。
		// 随后每次我们从栈顶取出一个节点 u，它是我们当前遍历到的节点，
		// 并把 u 的所有子节点从右向左逆序压入栈中，这样出栈的节点则是顺序从左向右的。
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return
}
