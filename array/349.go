package array

// 349. Intersection of Two Arrays
// 349. 两个数组的交集
// 思路：two set intersection
// time O(m+n) space O(m+n)
func intersection(nums1 []int, nums2 []int) (ans []int) {
	if len(nums1) == 0 || len(nums2) == 0 {
		return
	}
	s1, s2 := make(map[int]int, 0), make(map[int]int, 0)
	for _, num := range nums1 {
		s1[num]++
	}
	for _, num := range nums2 {
		s2[num]++
	}
	// 总是遍历较小的那个set
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	for key, _ := range s1 {
		if _, ok := s2[key]; ok {
			ans = append(ans, key)
		}
	}
	return
}
