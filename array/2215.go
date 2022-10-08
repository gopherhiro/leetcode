package array

// 2215. Find the Difference of Two Arrays
// 2215. 找出两数组的不同
// 思路：hashmap
// time O(m+n) space O(m+n)
func findDifference(nums1 []int, nums2 []int) [][]int {
	// 把两个数组的元素都存储哈希表
	h1, h2 := make(map[int]struct{}, 0), make(map[int]struct{}, 0)
	for _, v := range nums1 {
		h1[v] = struct{}{}
	}
	for _, v := range nums2 {
		h2[v] = struct{}{}
	}

	// 遍历两个哈希表，得到各自相对的差集
	// 哈希表中不存在重复元素，所以，得到的差集也就不存在重复元素
	ans := make([][]int, 2)
	for value := range h1 {
		if _, ok := h2[value]; ok {
			continue
		}
		ans[0] = append(ans[0], value)
	}
	for value := range h2 {
		if _, ok := h1[value]; ok {
			continue
		}
		ans[1] = append(ans[1], value)
	}

	return ans
}
