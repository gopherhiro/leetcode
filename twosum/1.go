package twosum

// 1. 两数之和
// 1. Two Sum
// 思路：hash + 遍历
// time O(N), space O(N)
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for k, v := range nums {
		if i, ok := hash[target-v]; ok {
			return []int{i, k}
		}
		hash[v] = k
	}
	return []int{-1, -1}
}
