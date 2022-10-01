package array

// 31. Next Permutation
// 31. 下一个排列
// 思路：Single Pass Approach
// time O(n) space O(1)
func nextPermutation(nums []int) {
	// 从后往前，找到第一个升序地址
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	// 存在下一个排列
	if i >= 0 {
		// 从后往前，找到第一个比「升序地址」大的数，交换
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	// 对「升序地址」后的数进行「逆置」
	reverse(nums, i+1)
}

func reverse(nums []int, start int) {
	i, j := start, len(nums)-1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
