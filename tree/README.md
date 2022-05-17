# 二叉树

```go
// 前序、中序遍历唯一确定一颗树。
// 中序、后序遍历唯一确定一颗树。
// 前序、后序遍历，不能唯一确定一颗树。
func main() {
/*	preorder := []int{2, 2, 3, 5, 5, 7}
	inorder := []int{2, 3, 2, 5, 5, 7}
	t := buildTreeFromPreIn(preorder, inorder)*/

//创建二叉树
t := New(7, 1)

// 输出
preorderTraverse(t)
fmt.Println()
inorderTraverse(t)
fmt.Println()
}

```

## 题目列表
- [94. 二叉树的中序遍历](94.go)
- [98. Validate Binary Search Tree](98.go)
- [102. Binary TreeNode Level Order Traversal](102.go)
- [103. Binary TreeNode Zigzag Level Order Traversal](103.go)
- [104. Maximum Depth of Binary TreeNode](104.go)
- [105. Construct Binary TreeNode from Preorder and Inorder Traversal](105.go)
- [106. Construct Binary TreeNode from Inorder and Postorder Traversal](106.go)
- [108. 将有序数组转换为二叉搜索树](108.go)
- [110. Balanced Binary Tree](110.go)
- [111. Minimum Depth of Binary TreeNode](111.go)
- [114. Flatten Binary TreeNode to Linked List](114.go)
- [116. 填充每个节点的下一个右侧节点指针](116.go)
- [124. Binary TreeNode Maximum Path Sum](124.go)
- [222. Count Complete TreeNode Nodes](222.go)
- [226. Invert Binary TreeNode](226.go)
- [230. Kth Smallest Element in a BST](230.go)
- [235. Lowest Common Ancestor of a Binary Search Tree](235.go)
- [236. Lowest Common Ancestor of a Binary TreeNode](236.go)
- [297. Serialize and Deserialize Binary TreeNode](297.go)
- [429. N-ary TreeNode Level Order Traversal](429.go)
- [450. Delete Node in a BST](450.go)
- [501. Find Mode in Binary Search TreeNode](501.go)
- [530. Minimum Absolute Difference in BST](530.go)
- [538. 把二叉搜索树转换为累加树](538.go)
- [543. Diameter of Binary TreeNode](543.go)
- [559. Maximum Depth of N-ary TreeNode](559.go)
- [589. N-ary TreeNode Preorder Traversal](589.go)
- [590. N-ary TreeNode Postorder Traversa](590.go)
- [606. Construct String from Binary TreeNode](606.go)
- [652. Find Duplicate Subtrees](652.go)
- [653. 两数之和 IV - 输入 BST](653.go)
- [654. Maximum Binary TreeNode](654.go)
- [700. Search in a Binary Search TreeNode](700.go)
- [701. Insert into a Binary Search TreeNode](701.go)
- [889. Construct Binary TreeNode from Preorder and Postorder Traversal](889.go)
- [958. 二叉树的完全性检验](958.go)
- [965. Univalued Binary TreeNode](965.go)
- [968. Binary TreeNode Cameras](968.go)