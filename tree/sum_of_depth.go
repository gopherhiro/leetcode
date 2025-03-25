package tree

import "fmt"

// 定义二叉树节点
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 递归解法
func sumOfDepths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	dfs(root, 0, &sum)
	return sum
}

func dfs(node *TreeNode, depth int, sum *int) {
	if node == nil {
		return
	}
	*sum += depth
	dfs(node.Left, depth+1, sum)
	dfs(node.Right, depth+1, sum)
}

// 迭代解法（BFS）
func sumOfDepthsBFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	queue := []*TreeNode{root}
	depth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += depth
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		depth++
	}
	return sum
}

// 构建示例树
func buildTree() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}
	return root
}

func main() {
	// 构建题目中的示例树
	tree := buildTree()

	// 运行递归解法
	resultDFS := sumOfDepths(tree)
	fmt.Println("递归解法 (DFS) 结果:", resultDFS)

	// 运行迭代解法
	resultBFS := sumOfDepthsBFS(tree)
	fmt.Println("迭代解法 (BFS) 结果:", resultBFS)
}
