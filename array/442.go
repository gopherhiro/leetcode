package array

// 442. 数组中重复的数据
// 策略：hashtable
// time O(N) space O(N)
func findDuplicates(nums []int) (ans []int) {
	if len(nums) == 0 {
		return ans
	}
	ht := make(map[int]int, 0)
	for _, num := range nums {
		ht[num]++
		if ht[num] == 2 {
			ans = append(ans, num)
		}
	}
	return ans
}
