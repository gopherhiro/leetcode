package hashtable

// 1773. Count Items Matching a Rule
// 1773. 统计匹配检索规则的物品数量
// 思路：哈希映射
// time O(n) space O(1)
func countMatches(items [][]string, key string, value string) int {
	// map rule to index of items
	dict := map[string]int{
		"type":  0,
		"color": 1,
		"name":  2,
	}
	// get the index of the key
	index := dict[key]

	// just count it
	count := 0
	for _, item := range items {
		if item[index] == value {
			count++
		}
	}
	return count
}
