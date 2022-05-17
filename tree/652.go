package tree

import "fmt"

// 652. Find Duplicate Subtrees
// time O(N)
// space O(N)
// 思路：二叉树的前序/中序/后序遍历结果可以描述二叉树的结构。
// 利用DFS后序遍历，序列化二叉树的Subtrees到hashtable，查找比对即可。
var res []*TreeNode
var hashmap = map[string]int{}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	findDupSubtreesHelper(root)
	return res
}

func findDupSubtreesHelper(root *TreeNode) string {
	// base case
	if root == nil {
		return "#"
	}

	left := findDupSubtreesHelper(root.Left)
	right := findDupSubtreesHelper(root.Right)

	// 序列化二叉树的子树
	rootValueStr := fmt.Sprintf("%d", root.Val)

	// 需要加逗号进行区分，以防止数字不同，但序列化字符串一致情况。
	subtree := left + "," + right + "," + rootValueStr

	// 把子树存储到map中，构建所有子树个数统计
	hashmap[subtree] += 1

	// 已经存在map中，则subtree存在重复
	// 对于同一类的重复子树，只需要返回其中任意一棵的根结点即可。
	// 计数为2表示有重复。
	if count, ok := hashmap[subtree]; ok {
		if count == 2 {
			res = append(res, root)
		}
	}

	return subtree
}
