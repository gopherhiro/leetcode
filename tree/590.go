package tree

// 590. N-ary TreeNode Postorder Traversa
// 思路：使用N叉树的后序遍历框架
// DFS: time O(N), space O(N)
func postorderOfNTree(root *Node) (answer []int) {
	var traverse func(root *Node)
	traverse = func(root *Node) {
		if root == nil {
			return
		}
		// 前序遍历位置
		for _, child := range root.Children {
			traverse(child)
		}
		// 后序遍历位置
		answer = append(answer, root.Val)
	}
	traverse(root)
	return
}

// Iteration : time O(N), space O(N)
// 如果看不懂，模拟一下。
func postorderOfNTreeIter(root *Node) (answer []int) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	vis := map[*Node]bool{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]

		// 如果当前节点为叶子节点 || 当前节点的子节点已经遍历过
		// 即这个条件，区分了N叉树前序遍历和后序遍历。
		if len(node.Children) == 0 || vis[node] {
			answer = append(answer, node.Val)
			stack = stack[:len(stack)-1]
			continue
		}
		// 首先把根节点入栈，因为根节点是前序遍历中的第一个节点。
		// 随后每次我们从栈顶取出一个节点 u，它是我们当前遍历到的节点，
		// 并把 u 的所有子节点从右向左逆序压入栈中，这样出栈的节点则是顺序从左向右的。
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		vis[node] = true
	}
	return
}

func postorderOfNTreeIterN(root *Node) (answer []int) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]

		// 如果当前节点为叶子节点 || 当前节点的子节点已经遍历过
		if len(node.Children) == 0 {
			answer = append(answer, node.Val)
			stack = stack[:len(stack)-1]
			continue
		}
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
		// 统一判断条件。与上面的解法对比，标记已经访问过该节点，可读性会更好。
		node.Children = nil
	}
	return
}
