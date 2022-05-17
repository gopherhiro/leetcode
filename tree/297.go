package tree

import (
	"fmt"
	"strconv"
	"strings"
)

// 297. Serialize and Deserialize Binary TreeNode
// 思路：前序遍历序列化二叉树，前序遍历反序列化二叉树即可。
// time O(N)
// space O(N)
type BinaryTree struct {
	root       *TreeNode
	serRes     []string
	deserInput []string
}

func Constructor() *BinaryTree {
	return &BinaryTree{root: nil, serRes: []string{}, deserInput: []string{}}
}

const NULL = "#"

// Serializes a tree to a single string.
func (bt *BinaryTree) serialize(root *TreeNode) string {
	bt.serializeHelper(root)
	return strings.Join(bt.serRes, ",")
}

func (bt *BinaryTree) serializeHelper(root *TreeNode) {
	// base case
	if root == nil {
		bt.serRes = append(bt.serRes, NULL)
		return
	}

	/****** 前序遍历位置 ******/
	// 收集树节点，即序列化
	rootValueStr := fmt.Sprintf("%d", root.Val)
	bt.serRes = append(bt.serRes, rootValueStr)

	// 递归调用，对左右子树序列化
	bt.serializeHelper(root.Left)
	bt.serializeHelper(root.Right)
}

// Deserializes your encoded data to tree.
func (bt *BinaryTree) deserialize(s string) *TreeNode {
	bt.deserInput = strings.Split(s, ",")
	return bt.deserializeHelper()
}

func (bt *BinaryTree) deserializeHelper() *TreeNode {
	/****** 前序遍历位置 ******/
	// 单单前序遍历结果是不能还原二叉树结构的
	// 因为缺少空指针的信息，至少要得到前、中、后序遍历中的两种才能还原二叉树。
	// 但是这里的 node 列表包含空指针的信息，所以只使用 node 列表就可以还原二叉树。
	// deserInput 这个必须是全局变量或传引用
	rootStr := bt.deserInput[0]
	bt.deserInput = bt.deserInput[1:]
	if rootStr == NULL {
		return nil
	}

	rootValue, _ := strconv.Atoi(rootStr)
	root := &TreeNode{Val: rootValue}

	root.Left = bt.deserializeHelper()
	root.Right = bt.deserializeHelper()

	return root
}
