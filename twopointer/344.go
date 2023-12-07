package twopointer

// 344. Reverse String
// 344. 反转字符串
// 思路：左右指针，交换元素即可。
// time O(N), space O(1)
func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStringN(str string) {
	if len(str) == 0 {
		return
	}
	s := []byte(str)
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// 反转数组
// 思路：左右指针
// time O(N) space O(1)
func reverseArray(nums []int) {
	if len(nums) == 0 {
		return
	}
	left, right := 0, len(nums)-1
	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}
