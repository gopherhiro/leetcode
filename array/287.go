package array

import "sort"

// 287. Find the Duplicate Number
// 287. 寻找重复数
// 思路：Sort + 遍历查找
func findDuplicate(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// 思路：遍历 + 哈希表
func findDuplicateN(nums []int) int {
	hash := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			return num
		}
		hash[num] = true
	}
	return -1
}

// 思路：Array as HashMap (Iterative)
func findDuplicateM(nums []int) int {
	for nums[0] != nums[nums[0]] {
		nums[0], nums[nums[0]] = nums[nums[0]], nums[0]
	}
	return nums[0]
}

// 思路：Cycle Detection，快慢指针
func findDuplicateO(nums []int) int {
	slow := nums[0]
	fast := nums[0]

	// 因为下面即将判断slow != fast
	// 所以，这里需要提前走，然后才能进入判断。
	slow = nums[slow]
	fast = nums[nums[fast]]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// Find the "entrance" to the cycle.
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
