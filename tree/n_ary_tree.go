package tree

import (
	"fmt"
)

// 定义 N 叉树节点
type Node struct {
	Val      int
	Children []*Node
}

// 输出 N 叉树第 n 层节点值
func printNaryTreeAtLevel(root *Node, level int) {
	if root == nil {
		return
	}

	if level == 1 {
		fmt.Print(root.Val, " ")
		return
	}

	for _, child := range root.Children {
		printNaryTreeAtLevel(child, level-1)
	}
}

// 获取 N 叉树的高度
func getHeight(root *Node) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	for _, child := range root.Children {
		childDepth := getHeight(child)
		if childDepth > maxDepth {
			maxDepth = childDepth
		}
	}

	return maxDepth + 1
}

func main() {
	// 构建示例 N 叉树
	root := &Node{Val: 1}
	child1 := &Node{Val: 2}
	child2 := &Node{Val: 3}
	child3 := &Node{Val: 4}
	child4 := &Node{Val: 5}
	child5 := &Node{Val: 6}

	root.Children = []*Node{child1, child2, child3}
	child1.Children = []*Node{child4, child5}

	n := 2 // 获取第 n 层节点值
	height := getHeight(root)

	if n <= height {
		fmt.Printf("Nodes at level %d: ", n)
		printNaryTreeAtLevel(root, n)
		fmt.Println()
	} else {
		fmt.Println("Invalid level")
	}
}
