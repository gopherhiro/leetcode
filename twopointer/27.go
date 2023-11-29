package twopointer

// 27. Remove Element
// 27. 移除元素
// 思路：two pointer
// time O(N) space O(1)
func removeElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] != val {
			// 维护 [0..slow-1] 不重复
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	// slow 即是去重后的数组长度
	return slow
}

// 27. Remove Element
// 27. 移除元素
// 思路：two pointer
// time O(N) space O(1)
func removeElementII(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow, fast := 0, 0
	for fast < n {
		if nums[fast] == val {
			fast++
			continue
		}
		// 维护 [0..slow-1] 不重复
		nums[slow] = nums[fast]
		slow++
		fast++
	}
	// slow 即是去重后的数组长度
	return slow
}
