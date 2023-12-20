package topk

import (
	"sort"
)

// 451. Sort Characters By Frequency
// 451. 根据字符出现频率排序
// 策略：桶排序
// time O(N) space O(N)
func frequencySort23(s string) string {
	// 统计频率
	ht := make(map[rune]int, 0)
	for _, v := range s {
		ht[v]++
	}
	// 放入桶中
	buckets := make([][]rune, len(s)+1)
	for ch, cnt := range ht {
		buckets[cnt] = append(buckets[cnt], ch)
	}

	// 构造频次字符
	ans := make([]rune, 0, len(s))
	for i := len(buckets) - 1; i >= 0; i-- {
		// 空桶，跳过
		if len(buckets[i]) == 0 {
			continue
		}
		// 使用桶中字符、频次，构造新字符串
		for j := 0; j < len(buckets[i]); j++ {
			ch := buckets[i][j]
			chs := make([]rune, 0, i)
			for t := 0; t < i; t++ {
				chs = append(chs, ch)
			}
			ans = append(ans, chs...)
		}
	}
	return string(ans)
}

// 451. Sort Characters By Frequency
// 451. 根据字符出现频率排序
// 策略：桶排序
// time O(N) space O(N)
func frequencySortN(s string) string {
	// 统计频率
	ht := make(map[rune]int, 0)
	for _, v := range s {
		ht[v]++
	}
	// 放入桶中
	buckets := make([][]rune, len(s)+1)
	for ch, cnt := range ht {
		buckets[cnt] = append(buckets[cnt], ch)
	}

	// 构造频次字符
	ans := make([]rune, 0, len(s))
	for i := len(buckets) - 1; i >= 0; i-- {
		// 空桶，跳过
		if len(buckets[i]) == 0 {
			continue
		}
		// 使用桶中字符、频次，构造新字符串
		for j := 0; j < len(buckets[i]); j++ {
			ch := buckets[i][j]
			chs := repeat([]rune{ch}, i)
			ans = append(ans, chs...)
		}
	}
	return string(ans)
}

// 451. Sort Characters By Frequency
// 451. 根据字符出现频率排序
// 策略：哈希表
// 思路：统计频率 + 排序
// time O(NlogN) space O(N)
func frequencySort(s string) string {
	// 统计频率
	ht := make(map[rune]int, 0)
	for _, v := range s {
		ht[v]++
	}
	// 排序
	type pair struct {
		char  rune
		count int
	}
	pairs := make([]pair, 0, len(ht))
	for ch, cnt := range ht {
		p := pair{
			char:  ch,
			count: cnt,
		}
		pairs = append(pairs, p)
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	// 取数据并还原字符串
	ans := make([]rune, 0, len(s))
	for _, p := range pairs {
		chars := repeat([]rune{p.char}, p.count)
		ans = append(ans, chars...)
	}
	return string(ans)
}

// source : bytes.Repeat
// repeat returns a new rune slice consisting of count copies of b.
func repeat(b []rune, count int) []rune {
	if count == 0 {
		return []rune{}
	}
	if count < 0 {
		panic("rune: negative Repeat count")
	} else if len(b)*count/count != len(b) {
		panic("rune: Repeat count causes overflow")
	}

	nb := make([]rune, len(b)*count)
	bp := copy(nb, b)
	for bp < len(nb) {
		copy(nb[bp:], nb[:bp])
		bp *= 2
	}
	return nb
}
