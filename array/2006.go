package array

// 2006. Count Number of Pairs With Absolute Difference K
// 2006. 差的绝对值为 K 的数对数目
// 思路：hashtable
// time O(n) space O(n)
func countKDifference(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	count := 0
	// |numi - numj| = k
	// numi - numj = +k or -k
	// so:
	// numi = numj + k
	// numi = numj - k
	for _, numj := range nums {
		// count number of numi
		count += m[numj+k]
		count += m[numj-k]
		m[numj]++
	}
	return count
}

// 2006. Count Number of Pairs With Absolute Difference K
// 2006. 差的绝对值为 K 的数对数目
// 思路：遍历
// time O(n^2) space O(1)
func countKDifference2(nums []int, k int) int {
	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(nums[i], nums[j]) == k {
				count++
			}
		}
	}
	return count
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
