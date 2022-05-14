package binarytree

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 116. 填充每个节点的下一个右侧节点指针
// Populating Next Right Pointers in Each Node
// time O(N) 遍历二叉树的节点
// space O(N) 函数递归调用栈
func connect(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left, root.Right)
	return root
}

func connectTwoNode(left *TreeNode, right *TreeNode) {
	if left == nil || right == nil {
		return
	}
	/**** 前序遍历位置 ****/
	// 将传入的两个节点连接
	left.Next = right

	// 连接相同父节点的两个子节点
	connectTwoNode(left.Left, left.Right)
	connectTwoNode(right.Left, right.Right)

	// 连接跨越父节点的两个子节点
	connectTwoNode(left.Right, right.Left)
}

// 116. 填充每个节点的下一个右侧节点指针
// Populating Next Right Pointers in Each Node
func addLinkToTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	addLink(root)
	return root
}
func addLink(root *TreeNode) {
	if root.Left == nil || root.Right == nil {
		return
	}

	/**** 前序遍历位置 ****/
	// 连接相同父节点的两个子节点
	root.Left.Next = root.Right

	// 连接跨越父节点的两个子节点
	if root.Next != nil {
		root.Right.Next = root.Next.Left
	}

	addLink(root.Left)
	addLink(root.Right)
}

// 114. Flatten Binary TreeNode to Linked List
// time O(N^2) 遍历的二叉树节点个数 O(N)，最后右子树的拼接 O(N)。
// space O(N) 取决于函数递归调用栈。
func flatten(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := flatten(root.Left)
	right := flatten(root.Right)

	/**** 后序遍历位置 ****/
	// 1、左右子树已经被拉平成一条链表
	// 2、将左子树作为右子树
	root.Left = nil
	root.Right = left

	// 3、将原先的右子树接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

	return root
}

// 226. Invert Binary TreeNode
// time O(N): N 为二叉树节点的数目
// space O(N): 使用的空间由递归栈的深度决定
// 思路：遍历二叉树的所有节点，并且对每个节点实施「交换子节点」的操作。
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	/*** 前序位置 ***/
	// 交换 root 节点的左右子节点
	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	// 左右子节点继续翻转它们的子节点
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// 方案2：分解问题解法。
// 思路：左子树的节点翻转，右子树的节点翻转，然后再交换左右子树。
func invertTreeN(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTreeN(root.Left)
	right := invertTreeN(root.Right)
	root.Left = right
	root.Right = left
	return root
}

// 654. Maximum Binary TreeNode
// 基本思路：先要找到根节点，然后让 build 函数递归生成左右子树即可。
func constructMaxBinaryTree(nums []int) *TreeNode {
	root := build(nums, 0, len(nums)-1)
	return root
}

func build(nums []int, begin, end int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if begin > end {
		return nil
	}

	// 找到数组中最大值及其索引
	maxValue, index := math.MinInt8, -1
	for i := begin; i <= end; i++ {
		if nums[i] > maxValue {
			maxValue = nums[i]
			index = i
		}
	}

	// 构建根结点
	root := &TreeNode{Val: maxValue}

	// 递归调用构建左右子树
	root.Left = build(nums, begin, index-1)
	root.Right = build(nums, index+1, end)

	return root
}

// 654. Maximum Binary TreeNode
// 基本思路：先要找到根节点，然后让 build 函数递归生成左右子树即可。
// time O(N^2)
// space O(N)
func constructMaximumBinaryTree(nums []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	// 找到数组中最大值的索引，构建二叉树节点
	index := findMaxValueIndex(nums)
	root := &TreeNode{Val: nums[index]}

	// 递归调用构建左右子树
	root.Left = constructMaximumBinaryTree(nums[:index])
	root.Right = constructMaximumBinaryTree(nums[index+1:])
	return root
}

func findMaxValueIndex(nums []int) int {
	maxValueIndex := 0
	for i, v := range nums {
		if v > nums[maxValueIndex] {
			maxValueIndex = i
		}
	}
	return maxValueIndex
}

// 105. Construct Binary TreeNode from Preorder and Inorder Traversal
// 思路：构造二叉树，第一件事一定是找根节点，然后想办法构造左右子树。
// 通过前序或者后序遍历结果找到根节点，然后根据中序遍历结果确定左右子树。
// time O(N^2)
// space O(N)
func buildTreeFromPreIn(preorder []int, inorder []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// root 节点对应的值就是前序遍历数组的第一个元素
	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	// rootValue 在中序遍历数组中的索引
	inorderIndex := findInorderIndex(inorder, rootValue)

	// 递归调用构建左右子树 : 这里的下标需要计算
	root.Left = buildTreeFromPreIn(preorder[1:inorderIndex+1], inorder[:inorderIndex])
	root.Right = buildTreeFromPreIn(preorder[inorderIndex+1:], inorder[inorderIndex+1:])
	return root
}

func findInorderIndex(inorder []int, value int) int {
	for i, v := range inorder {
		if v == value {
			return i
		}
	}
	return 0
}

// 106. Construct Binary TreeNode from Inorder and Postorder Traversal
// 思路：构造二叉树，第一件事一定是找根节点，然后想办法构造左右子树。
// 通过前序或者后序遍历结果找到根节点，然后根据中序遍历结果确定左右子树。
// time O(N^2)
// space O(N)
func buildTreeFromInPost(inorder []int, postorder []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	if len(postorder) == 1 {
		return &TreeNode{Val: postorder[0]}
	}

	// root 节点对应的值就是后序遍历数组的最后一个元素
	rootValue := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootValue}

	// rootValue 在中序遍历数组中的索引
	inorderIndex := findInorderIndex(inorder, rootValue)

	root.Left = buildTreeFromInPost(inorder[:inorderIndex], postorder[:inorderIndex])
	root.Right = buildTreeFromInPost(inorder[inorderIndex+1:], postorder[inorderIndex:len(postorder)-1])

	return root
}

// 889. Construct Binary TreeNode from Preorder and Postorder Traversal
// 思路：构造二叉树，第一件事一定是找根节点，然后想办法构造左右子树。
// 通过前序或者后序遍历结果找到根节点，然后根据中序遍历结果确定左右子树。
// time O(N^2)
// space O(N)
func buildTreeFromPrePost(preorder, postorder []int) *TreeNode {
	/*** 前序遍历位置 ***/
	// base case
	if len(preorder) == 0 || len(postorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	// root 节点对应的值就是前序遍历数组的第一个元素
	rootValue := preorder[0]
	root := &TreeNode{Val: rootValue}

	// 把前序遍历结果的第二个元素作为左子树的根节点的值。
	leftRootValue := preorder[1]
	index := rootLeftIndex(postorder, leftRootValue)

	// 在后序遍历结果中寻找左子树根节点的值，从而确定了左子树的索引边界，进而确定右子树的索引边界
	// 从而递归构造左右子树即可。
	root.Left = buildTreeFromPrePost(preorder[1:index+1], postorder[:index+1])
	root.Right = buildTreeFromPrePost(preorder[index+1:], postorder[index:len(postorder)-1])
	return root
}

func rootLeftIndex(postorder []int, value int) int {
	for i, v := range postorder {
		if v == value {
			// 这里返回的索引需要+1
			return i + 1
		}
	}
	return 0
}

// 652. Find Duplicate Subtrees
// time O(N)
// space O(N)
// 思路：二叉树的前序/中序/后序遍历结果可以描述二叉树的结构。
// 利用DFS后序遍历，序列化二叉树的Subtrees到hashtable，查找比对即可。
var res []*TreeNode
var hashmap = map[string]int{}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	findDupSubtreesHelper(root)
	fmt.Println(hashmap)
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

// 102. Binary TreeNode Level Order Traversal
// time O(N)
// space O(N)
func levelOrderTraverse(t *TreeNode) [][]int {
	if t == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// 根节点入队
	queue = append(queue, t)

	// 第一层循环控制：从上向下层层遍历
	for len(queue) > 0 {
		// 记录这一层的节点值
		row := make([]int, 0)
		rowSize := len(queue)

		// 第二层循环控制每一层：从左向右遍历
		for i := 0; i < rowSize; i++ {
			// 从队列中弹出一个节点处理
			cur := queue[0]
			queue = queue[1:]
			row = append(row, cur.Val)

			// 其有左右子节点，则左右子节点入队
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}

		res = append(res, row)
	}

	return res
}

// 103. Binary TreeNode Zigzag Level Order Traversal
// 思路：与二叉树层序遍历一样，只要再加一个bool变量控制遍历方向即可。
// time O(N), space O(N)
func zigzagLevelOrder(t *TreeNode) [][]int {
	if t == nil {
		return [][]int{}
	}

	res := make([][]int, 0)
	queue := make([]*TreeNode, 0)

	// 方向标识：向右
	toRight := true

	// 根节点入队
	queue = append(queue, t)

	// 第一层循环控制：从上向下层层遍历
	for len(queue) > 0 {
		// 记录这一层的节点值
		row := make([]int, 0)
		rowSize := len(queue)

		// 第二层循环控制每一层：从左向右遍历
		for i := 0; i < rowSize; i++ {
			// 从队列中弹出一个节点处理
			cur := queue[0]
			queue = queue[1:]

			// 实现之字遍历
			if toRight {
				row = append(row, cur.Val)
			} else {
				row = append([]int{cur.Val}, row...)
			}

			// 其有左右子节点，则左右子节点入队
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 切换方向
		toRight = !toRight
		res = append(res, row)
	}

	return res
}

// 94. Binary TreeNode Inorder Traversal
// 递归方案
// time O(N)
// space O(N)
type InorderTraversal struct {
	inorderRes []int
}

func NewInorder() *InorderTraversal {
	return &InorderTraversal{inorderRes: []int{}}
}

func (it *InorderTraversal) inorderTraversal(root *TreeNode) []int {
	it.inorderTraversalHelper(root)
	return it.inorderRes
}

func (it *InorderTraversal) inorderTraversalHelper(root *TreeNode) {
	if root == nil {
		return
	}
	it.inorderTraversalHelper(root.Left)
	it.inorderRes = append(it.inorderRes, root.Val)
	it.inorderTraversalHelper(root.Right)
}

// 迭代方案：使用栈结构，模拟函数递归调用栈。
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			// 取栈顶元素并出栈
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			root = node.Right
		}

	}
	return res
}

// 分解问题
// 动态规划方案：牛！
func inorderTraversN(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	left := inorderTraversN(root.Left)
	res = append(res, left...)

	res = append(res, root.Val)

	right := inorderTraversN(root.Right)
	res = append(res, right...)

	return res
}

func preorderTraversN(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	left := preorderTraversN(root.Left)
	res = append(res, left...)

	right := preorderTraversN(root.Right)
	res = append(res, right...)

	return res
}

func postorderTraversN(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	left := postorderTraversN(root.Left)
	res = append(res, left...)

	right := postorderTraversN(root.Right)
	res = append(res, right...)

	res = append(res, root.Val)

	return res
}

// 104. Maximum Depth of Binary TreeNode
// 定义：输入一个节点，返回以该节点为根的二叉树的最大深度
// 动态规划思路：根据左右子树的最大深度推出原二叉树的最大深度即可。
// time O(N), space O(N)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	// 根据左右子树的最大深度推出原二叉树的最大深度
	return maxValue(leftDepth, rightDepth) + 1
}

func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 543. Diameter of Binary TreeNode
// 思路：所谓二叉树的直径，就是左右子树的最大深度之和，
// 那么直接的想法是对每个节点计算左右子树的最大高度，
// 得出每个节点的直径，从而得出最大的那个直径。
// time O(N), space O(N)
func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int

	var maxDepth func(*TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDepth(root.Left)
		right := maxDepth(root.Right)
		// 二叉树的直径：左右子树的最大深度之和，比较&获取最大直径。
		maxDiameter = maxValue(maxDiameter, left+right)
		return maxValue(left, right) + 1
	}
	maxDepth(root)

	return maxDiameter
}

// 111. Minimum Depth of Binary TreeNode
// 思路：BFS（Breadth-First Search）
// 从上到下，从左至右，第一个达到叶子节点的层，即是树的最小深度。
// BFS与DFS的区别，BFS第一次搜索到的结果是最优的。
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[0]
			queue = queue[1:]
			// 判断是否到达叶子结点
			if cur.Left == nil && cur.Right == nil {
				return depth
			}

			// 层序遍历的框架
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
		}
		// 更新深度
		depth++
	}
	return depth
}

// 124. Binary TreeNode Maximum Path Sum
// 定义：给你一个二叉树的根节点 root ，返回其 最大路径和。
// 路径和 是路径中各节点值的总和。
// 递归-分解问题：time O(N), space O(N)
var maxPathSumRes = math.MinInt16

func maxPathSum(root *TreeNode) int {
	maxPathSumOneSide(root)
	return maxPathSumRes
}

func maxPathSumOneSide(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// 取得左右子树的最大单边路径和。
	left := maxValue(maxPathSumOneSide(root.Left), 0)
	right := maxValue(maxPathSumOneSide(root.Right), 0)

	// 后序遍历位置，顺便更新最大路径和
	sum := left + root.Val + right
	maxPathSumRes = maxValue(maxPathSumRes, sum)

	// 实现函数定义，左右子树的最大单边路径和加上根节点的值
	// 就是从根节点 root 为起点的最大单边路径和
	return maxValue(left, right) + root.Val
}

// 222. Count Complete TreeNode Nodes
// 思路：遍历完全二叉树，统计节点个数即可。
// time O(N), space O(N)
var treeNodesNumber = 0

func countTreeNodes(root *TreeNode) int {
	countNodesHelper(root)
	return treeNodesNumber
}

func countNodesHelper(root *TreeNode) {
	if root == nil {
		return
	}
	treeNodesNumber++
	countNodesHelper(root.Left)
	countNodesHelper(root.Right)
}

// 222. Count Complete TreeNode Nodes
// 思路：满二叉树节点计算公式+递归。
// time (logN^2), space O(N)
func countTreeNodesOptimal(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 记录左、右子树的高度
	left, right := root, root
	leftHigh, rightHigh := 0, 0
	for left != nil {
		left = left.Left
		leftHigh++
	}
	for right != nil {
		right = right.Right
		rightHigh++
	}
	// 如果左右子树的高度相同，则是一棵满二叉树
	// 满二叉树的节点个数，可以直接通过树高算出来。
	if leftHigh == rightHigh {
		nodesNumber := math.Pow(float64(2), float64(leftHigh)) - 1
		return int(nodesNumber)
	}
	// 如果左右高度不同，则按照普通二叉树的逻辑计算
	return countTreeNodesOptimal(root.Left) + countTreeNodesOptimal(root.Right) + 1
}

// 236. Lowest Common Ancestor of a Binary TreeNode
// time O(N), space O(N)
// 思路：遍历树，寻找p，q节点，从而确定最近公共祖先。
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// 后序位置
	// case 1
	// 如果 p 和 q 都在以 root 为根的树中，
	// 那么 left 和 right 一定分别是 p 和 q（从 base case 看出来的）
	if left != nil && right != nil {
		return root
	}

	// case 2
	// 如果 p 和 q 都不在以 root 为根的树中，直接返回 nil。
	if left == nil && right == nil {
		return nil
	}

	// case 3
	// 如果 p 和 q 只有一个存在于 root 为根的树中，函数返回该节点。
	if left == nil {
		return right
	}
	return left
}

// 235. Lowest Common Ancestor of a Binary Search Tree
// 235. 二叉搜索树的最近公共祖先
// time O(logN) space O(N)
func lowestCommonAncestorII(root, p, q *TreeNode) *TreeNode {
	// base case
	if root == nil {
		return nil
	}

	if p.Val > q.Val {
		// 保证 p <= q，方便下面二分。
		return lowestCommonAncestorII(root, q, p)
	}

	// p <= root <= q
	if root.Val >= p.Val && root.Val <= q.Val {
		return root
	}

	// p,q < root
	if q.Val < root.Val {
		return lowestCommonAncestorII(root.Left, p, q)
	}
	// p,q > root
	return lowestCommonAncestorII(root.Right, p, q)
}

// 501. Find Mode in Binary Search TreeNode
// 501. 二叉搜索树中的众数
// dfs inorder traverse + 策略逻辑
// time O(N), space O(N)
// 如果看不懂，那就模拟一下。
func findMode(root *TreeNode) []int {
	var base, count, maxCount int
	answer := make([]int, 0)

	// 一棵二叉搜索树(BST, binary search tree)的中序遍历序列是一个递增的有序序列。
	// 顺序扫描中序遍历序列：
	// 用 base 记录当前的数字，
	// 用 count 记录当前数字重复的次数，
	// 用 maxCount 来维护已经扫描过的数当中出现最多的那个数字的出现次数，
	// 用 answer 数组记录出现的众数。
	update := func(value int) {
		if value == base {
			count++
		} else {
			base, count = value, 1
		}

		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			maxCount = count
			answer = []int{base}
		}
	}

	// 考虑不存储中序遍历序列：
	// 如果在递归中序遍历的过程中访问节点，
	// 就可以省去存储中序遍历序列的空间。
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		update(root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return answer
}

// 559. Maximum Depth of N-ary TreeNode
// 思路：使用N叉树的遍历框架
// 深度优先遍历（DFS）：time O(N) space O(N)
// Definition for a Node.
type Node struct {
	Val      int
	Children []*Node
}

func maxDepthOfNTree(root *Node) int {
	// base case
	if root == nil {
		return 0
	}

	// 遍历N叉树，得到子树的最大深度
	var maxChildDepth int
	for _, child := range root.Children {
		childDepth := maxDepthOfNTree(child)
		if childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}

	// N 个子树的最大深度中的最大值为 maxChildDepth
	// 则该 N 叉树的最大深度为 maxChildDepth+1。
	return maxChildDepth + 1
}

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

// 429. N-ary TreeNode Level Order Traversal
// 思路：参考二叉树层序遍历。
// 使用队列：time O(N), space O(N)
func levelOrderNTree(root *Node) (answer [][]int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		row := make([]int, 0)
		rowSize := len(queue)
		for i := 0; i < rowSize; i++ {
			cur := queue[0]
			queue = queue[1:]
			row = append(row, cur.Val)
			if cur.Children != nil {
				queue = append(queue, cur.Children...)
			}
		}
		answer = append(answer, row)
	}
	return
}

// 968. Binary TreeNode Cameras
// DFS post order traversal, time O(N), space O(N)
// 思路：后序位置可以接收到子树的信息，同时也可以通过函数参数接收到父节点传递的信息。
func minCameraCover(root *TreeNode) (answer int) {
	if root == nil {
		answer = 0
	}

	const NodeIsNIL = -1 // 空节点
	const (
		Uncover = iota // 监控未覆盖
		Cover          // 监控已覆盖
		Set            // 安装摄像头
	)

	var setCamera func(root *TreeNode, hasParent bool) int
	setCamera = func(root *TreeNode, hasParent bool) int {
		// base case
		if root == nil {
			return NodeIsNIL
		}
		// 获取左右子节点的情况
		left := setCamera(root.Left, true)
		right := setCamera(root.Right, true)

		// 根据左右子节点的情况和父节点的情况
		// 判断当前节点应该做的事情
		if left == NodeIsNIL && right == NodeIsNIL {
			// 当前节点是叶子节点
			if hasParent {
				// 有父节点的话，让父节点来 Cover 自己
				return Uncover
			}
			// 没有父节点的话，自己 Set （安装）一个摄像头
			answer++
			return Set
		}

		// 左右子树存在没有被 cover 的
		// 必须在当前节点 set 一个摄像头
		if left == Uncover || right == Uncover {
			answer++
			return Set
		}

		// 左右子树只要有一个 Set 了摄像头
		// 当前节点就已经是 Cover 状态了
		if left == Set || right == Set {
			return Cover
		}

		// 最后一种情况：left == Cover && right == Cover
		// 即当前节点的左右子节点都被 Cover
		if hasParent {
			// 有父节点，可以等父节点 Cover 自己
			return Uncover
		}
		// 没有父节点，只能自己 Set 一个摄像头
		answer++
		return Set
	}

	setCamera(root, false)
	return
}

// 965. Univalued Binary TreeNode
// Approach 1: Depth-First Search
// time O(N), space O(N)
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	isUnival := true

	// 记下第一个节点的值
	// 遍历二叉树，看看是否存在其他值。
	first := root.Val
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		// base case
		if root == nil {
			return
		}
		if first != root.Val {
			isUnival = false
			return
		}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)

	return isUnival
}

// Approach 2: Recursion
// 分解问题：左右子树都是单值的，则该树是单值的。
// time O(N), space O(Height)
func isUnivalTreeRecur(root *TreeNode) bool {
	// 左子树是单值的
	leftIsUnival := root.Left == nil || root.Val == root.Left.Val && isUnivalTreeRecur(root.Left)
	// 右子树是单值的
	rightIsUnival := root.Right == nil || root.Val == root.Right.Val && isUnivalTreeRecur(root.Right)
	// 从而该二叉树是单值的。
	return leftIsUnival && rightIsUnival
}

// 606. Construct String from Binary TreeNode
// [root,#,#] => root
// [root,#,right] => root()(right)
// [root,left,#] => root(left)
// [root,left,right] => root(left)(right)
// Approach: Recursion, 分解问题: time O(N), space O(N)
func tree2str(root *TreeNode) string {
	// base case
	if root == nil {
		return ""
	}
	// [root,#,#] => root
	if root.Left == nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val) + ""
	}

	leftStr := tree2str(root.Left)
	rightStr := tree2str(root.Right)
	// 后序位置
	// [root,#,right] => root()(right)
	if root.Left == nil && root.Right != nil {
		return fmt.Sprintf("%d", root.Val) + "()" + "(" + rightStr + ")"
	}

	// [root,left,#] => root(left)
	if root.Left != nil && root.Right == nil {
		return fmt.Sprintf("%d", root.Val) + "(" + leftStr + ")"
	}

	// [root,left,right] => root(left)(right)
	return fmt.Sprintf("%d", root.Val) + "(" + leftStr + ")" + "(" + rightStr + ")"
}

// 230. Kth Smallest Element in a BST
// DFS-Recursive Inorder Traversal，time O(N), space O(N)
// 思路：BST，升序排序，然后找第 k 个元素。
func kthSmallest(root *TreeNode, k int) int {
	// 记录当前元素的排名
	var rank int
	// 记录结果
	var answer int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Left)
		/* 中序位置 */
		// 找到第k小的元素
		rank++
		if rank == k {
			answer = root.Val
			return
		}
		traverse(root.Right)
	}
	traverse(root)
	return answer
}

// Approach 2: Iterative Inorder Traversal
// The above recursion could be converted into iteration, with the help of stack.
// This way one could speed up the solution
// because there is no need to build the entire inorder traversal,
// and one could stop after the kth element.
// time (H+k), space (H)
func kthSmallestIter(root *TreeNode, k int) int {
	stack := make([]*TreeNode, 0)
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 中序位置
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--
		// 找到第k小的元素
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// 538. Convert BST to Greater TreeNode
// 538. 把二叉搜索树转换为累加树
// Approach #1 Recursion, time O(N), space O(N)
// 思路：采用中序降序遍历，可以得到该节点右边的所有节点值的和。
// 从而的到Greater Sum TreeNode
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var sum int
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		traverse(root.Right)

		// 维护累加和 & 更新节点值
		sum += root.Val
		root.Val = sum

		traverse(root.Left)
	}
	traverse(root)

	return root
}

// 700. Search in a Binary Search TreeNode
// Recursion, time O(H), space O(H)
func searchBST(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	// 去左子树搜索
	if target < root.Val {
		return searchBST(root.Left, target)
	}
	// 去右子树搜索
	if target > root.Val {
		return searchBST(root.Right, target)
	}
	return root
}

// 700. Search in a Binary Search TreeNode
// Iteration, time O(N), space O(1)
func searchBSTIter(root *TreeNode, target int) *TreeNode {
	for root != nil {
		if root.Val == target {
			return root
		}
		// 特别注意这里：要么走左子树，要么走右子树。
		if target < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}

// 701. Insert into a Binary Search TreeNode
// Recursion: time O(H), space O(H)
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到位置，插入新节点
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 如果插入值小于根节点，插入到左子树
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	}
	// 如果插入值，大于根节点，插入到右子树
	if val > root.Val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 108. Convert Sorted Array to Binary Search Tree
// 108. 将有序数组转换为二叉搜索树
// 思路：前序遍历构建二叉搜索树，中序结果保证高度平衡二叉搜索树。
// 一个有序数组对于 BST 来说就是中序遍历结果。
// 根节点在数组中心，数组左侧是左子树元素，右侧是右子树元素。
// time O(N), space O(N)
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	var buildBST func(left, right int) *TreeNode
	buildBST = func(left, right int) *TreeNode {
		// base case
		if left > right {
			// 空区间
			return nil
		}

		// 前序位置

		// 构造根节点
		// BST 节点左小右大，中间的元素就是根节点
		mid := left + (right-left)>>1
		root := &TreeNode{Val: nums[mid]}

		// 递归构建左、右子树
		root.Left = buildBST(left, mid-1)
		root.Right = buildBST(mid+1, right)

		return root
	}
	return buildBST(0, len(nums)-1)
}

// 450. Delete Node in a BST
// Recursion, time O(H), space O(H)
// 删除时，注意区分情况：case 1，case 2，case 3，对应不同的删除操作。
func deleteNodeBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	// 找到了要删除的节点
	if val == root.Val {
		// case 1，要删除的节点是叶子节点（case 1可以删除，因为case 2情况已包含该情况）
		// 为了可读性，这里可以保留。
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// case 2，要删除的节点只有一个子节点
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		// case 3，要删除的节点有两个子节点
		// root.left != nil && root.right != nil
		// 找到右子树的最小节点
		minNode := getMin(root.Right)
		// 删除右子树最小的节点
		root.Right = deleteNodeBST(root.Right, minNode.Val)
		// 用右子树最小的节点替换 root 节点
		minNode.Left = root.Left
		minNode.Right = root.Right
		root = minNode
	}

	// 到左子树中删除
	if val < root.Val {
		root.Left = deleteNodeBST(root.Left, val)
	}

	// 到右子树中删除
	if val > root.Val {
		root.Right = deleteNodeBST(root.Right, val)
	}

	return root
}

func getMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// BST 最左边的就是最小的。
	for root.Left != nil {
		root = root.Left
	}
	return root
}

// 98. Validate Binary Search Tree
// DFS-post order traversal：Recursion, time O(N), space O(N)
// 思路：递归验证根节点是否在相应的范围内
// root.val => (min, max)
// root.left.val => (min, root.val)
// root.right.val => (root.val, max)
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var traverse func(root *TreeNode, min, max int) bool
	traverse = func(root *TreeNode, min, max int) bool {
		// base case
		if root == nil {
			return true
		}
		// root的值，必须满足：min < root.Val < max
		if root.Val <= min || root.Val >= max {
			return false
		}

		// 左子树的最大值是 root.val
		leftIsBST := traverse(root.Left, min, root.Val)

		// 右子树的最小值是 root.val
		rightIsBST := traverse(root.Right, root.Val, max)

		// left is BST and right is BST, so the tree is BST
		return leftIsBST && rightIsBST
	}
	return traverse(root, math.MinInt64, math.MaxInt64)
}

// 110. Balanced Binary Tree
// DFS post order traverse, time O(N), space O(N)
func isBalanced(root *TreeNode) bool {
	answer := true

	var maxDepth func(root *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftMaxDepth := maxDepth(root.Left)
		rightMaxDepth := maxDepth(root.Right)
		// 后序遍历位置
		// 如果左右子树最大深度大于 1，就不是平衡二叉树
		if absv(leftMaxDepth-rightMaxDepth) > 1 {
			answer = false
		}
		return 1 + maxValue(leftMaxDepth, rightMaxDepth)
	}

	maxDepth(root)

	return answer
}

func absv(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

// 530. Minimum Absolute Difference in BST
// DFS-inorder traverse, time O(N), space O(N)
func getMinimumDifference(root *TreeNode) int {
	answer := math.MaxInt32
	var prev *TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		// 中序位置，进行处理
		// 如果上一个节点存在，则进行最小差值更新。
		// 否则，只对上一个节点进行记录。
		if prev != nil {
			answer = min(answer, root.Val-prev.Val)
		}
		prev = root // 记录上一个节点
		inorder(root.Right)
	}
	inorder(root)
	return answer
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 653. 两数之和 IV - 输入 BST
// 653. Two Sum IV - Input is a BST
// 思路：遍历二叉树 + 哈希表
func findTarget(root *TreeNode, k int) (answer bool) {
	if root == nil {
		answer = false
	}
	hashtable := map[int]struct{}{}
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		if _, ok := hashtable[k-root.Val]; ok {
			answer = true
			return
		}
		hashtable[root.Val] = struct{}{}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return
}

// 前序、中序遍历唯一确定一颗树。
// 中序、后序遍历唯一确定一颗树。
// 前序、后序遍历，不能唯一确定一颗树。
func main() {
	/*	preorder := []int{2, 2, 3, 5, 5, 7}
		inorder := []int{2, 3, 2, 5, 5, 7}
		t := buildTreeFromPreIn(preorder, inorder)*/

	//创建二叉树
	//t := New(7, 1)

	// 输出
	/*	preorderTraverse(t)
		fmt.Println()
		inorderTraverse(t)
		fmt.Println()*/

	fmt.Println("res:")
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)

	preorderTraverse(res)
	fmt.Println()
	inorderTraverse(res)
	fmt.Println()
}
