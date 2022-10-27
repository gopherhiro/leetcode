package array

// 2176. Count Equal and Divisible Pairs in an Array
// 2176. 统计数组中相等且可以被整除的数对
// 思路：hashtable
// time O(n*m) space(n)
func countPairs(nums []int, k int) int {
	count := 0
	m := make(map[int][]int, 0)
	for i, v := range nums {
		// 元素不存在，直接将「元素和其索引」放到哈希表中
		if _, ok := m[v]; !ok {
			m[v] = append(m[v], i)
			continue
		}
		// 元素已存在，则遍历「记录的索引列表」进行对数统计
		// 再将「元素和其索引」放到哈希表中
		for _, j := range m[v] {
			if i*j%k == 0 {
				count++
			}
		}
		m[v] = append(m[v], i)
	}
	return count
}
