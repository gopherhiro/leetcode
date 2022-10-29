package array

// 2364. Count Number of Bad Pairs
// 2364. 统计坏数对的数目
// 思路：哈希表
// time O(n) space O(n)
func countBadPairs(nums []int) int64 {
	n := len(nums)
	count := 0
	m := make(map[int]int, 0)
	for i := 0; i < n; i++ {
		v := nums[i] - i
		// 查询时，v 表示 numj - j
		count += m[v]
		// 存储时，v 表示 numi - i
		m[v]++
	}
	// bad pairs = total pairs - nice pairs
	sn := n * (n - 1) / 2
	ans := sn - count
	return int64(ans)
}
