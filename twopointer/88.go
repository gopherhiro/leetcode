package twopointer

// 88. Merge Sorted Array
// 88. 合并两个有序数组
// 思路：two pointer
// time O(m+n) space O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		// 判断是否到达 num1 末尾
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		// 判断是否到达 num2 末尾
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}

		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}
