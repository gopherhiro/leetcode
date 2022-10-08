package array

import "sort"

// 350. Intersection of Two Arrays II
// 350. 两个数组的交集 II
// 思路：two pointer
// time O(nlogn + mlogm) space O(min(n, m))
func intersect2(nums1 []int, nums2 []int) []int {
	// 排序
	sort.Ints(nums1)
	sort.Ints(nums2)

	// 双指针遍历，收集交集。
	// 元素相等：发现交集元素，一起往后走
	// 元素不等：元素较小的数组指针往后走
	ans := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			ans = append(ans, nums1[i])
			i++
			j++
		}
	}
	return ans
}

// 350. Intersection of Two Arrays II
// 350. 两个数组的交集 II
// 思路：hashmap
// time O(n + m) space O(min(n, m))
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	// 统计较小数组元素频次
	counter := make(map[int]int, 0)
	for _, v := range nums1 {
		counter[v]++
	}
	// 遍历较大数组，收集交集元素
	ans := make([]int, 0)
	for _, v := range nums2 {
		if counter[v] > 0 {
			ans = append(ans, v)
			counter[v]--
		}
	}
	return ans
}
