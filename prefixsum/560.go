package prefixsum

// 560. Subarray Sum Equals K
// 560. 和为 K 的子数组
// 思路：前缀和+hash
// time O(N) space O(N)
func subarraySum1(nums []int, k int) int {
	hash := make(map[int]int, 0)
	hash[0] = 1
	ans, sum0i := 0, 0
	for i := 0; i < len(nums); i++ {
		sum0i += nums[i]   // sum[0..i]
		sum0j := sum0i - k // sum[0..j]
		if count, ok := hash[sum0j]; ok {
			ans += count
		}
		// sum[0..i]加入hash，并记录出现的次数。
		hash[sum0i] += 1
	}
	return ans
}

// 560. Subarray Sum Equals K
// 560. 和为 K 的子数组
// 思路：前缀和+穷举
// time O(N^2) space O(N)
func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	// 构造前缀和
	preSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}

	// 穷举所有子数组
	ans := 0
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			// nums[j..i-1]
			if preSum[i]-preSum[j] == k {
				ans++
			}
		}
	}

	return ans
}
