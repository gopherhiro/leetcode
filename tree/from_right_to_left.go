package tree

import "fmt"

// 输出二叉树第 n 层节点值（从右到左）
func printRightToLeftAtLevel(root *TreeNode, level int) {
	if root == nil {
		return
	}

	// 当前层为目标层级时输出节点值
	if level == 1 {
		fmt.Print(root.Val, " ")
		return
	}

	// 递归遍历左子树和右子树（从右到左输出）
	printRightToLeftAtLevel(root.Right, level-1)
	printRightToLeftAtLevel(root.Left, level-1)
}

// 获取二叉树的高度
func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	return max(leftHeight, rightHeight) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// 构建示例二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	n := 3 // 获取第 n 层节点值
	height := getHeight(root)

	if n <= height {
		fmt.Printf("Nodes at level %d (from right to left): ", n)
		printRightToLeftAtLevel(root, n)
		fmt.Println()
	} else {
		fmt.Println("Invalid level")
	}
}
