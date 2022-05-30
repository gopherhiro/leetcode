package array

// 347. Top K Frequent Elements
// 347. 前 K 个高频元素
// 思路：哈希表 + 桶排序
// time O(N) space O(N)
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	ht := make(map[int]int, 0)
	for _, v := range nums {
		ht[v]++
	}
	// 使用「桶排序」来进行频次筛选
	buckets := make([][]int, len(nums)+1)
	for num, cnt := range ht {
		buckets[cnt] = append(buckets[cnt], num)
	}

	ans := make([]int, 0)
	for i := len(buckets) - 1; i >= 0; i-- {
		// 空桶，跳过
		if len(buckets[i]) == 0 {
			continue
		}
		ans = append(ans, buckets[i]...)
		// 已经获得 top k，则停止筛选。
		if len(ans) == k {
			break
		}
	}

	return ans
}
