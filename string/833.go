package string

import "sort"

// 833. Find And Replace in String
// 833. 字符串中的查找与替换
// 思路：模拟，比较 & 替换
// time O(nm) space O(n)
func findReplaceString(s string, index []int, sources []string, targets []string) string {
	// 索引与索引下标的映射关系
	// 方便查询 sources, targets
	pairs := make([][]int, 0)
	for i := 0; i < len(index); i++ {
		pair := []int{index[i], i}
		pairs = append(pairs, pair)
	}

	// 倒序排序索引对，so we can replace s from right to left.
	// 这样就不会弄乱 s 中的下标顺序了。
	sort.Slice(pairs, func(i int, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	// verify equality of subring in s string and sources.
	// replace s string
	for _, v := range pairs {
		i := v[0] // i point to the index of s string
		p := v[1] // p point to the index of sources,targets
		src := sources[p]
		trg := targets[p]

		// 如果截取的子串超出 s 范围，
		// 则定不是 src 串，跳过即可。
		if i+len(src) > len(s) {
			continue
		}
		if s[i:i+len(src)] == src {
			s = s[:i] + trg + s[i+len(src):]
		}
	}

	return s
}
