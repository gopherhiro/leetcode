package tree

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
