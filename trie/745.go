package trie

// 745. Prefix and Suffix Search
// 745. 前缀和后缀搜索
// 思路：Trie
type WordFilter struct {
	children [26]*WordFilter
	isEnd    bool
	prefix   []int // prefix index
	suffix   []int // suffix index
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{}
	for index, word := range words {
		wf.insert(index, word)
	}
	return wf
}

// time O(n) space O(n)
func (wf *WordFilter) insert(index int, word string) {
	wf.prefixInsert(index, word)
	wf.suffixInsert(index, word)
}

// from left to right insert the character
func (wf *WordFilter) prefixInsert(index int, word string) {
	cur := wf
	for _, ch := range word {
		key := ch - 'a'
		if cur.children[key] == nil {
			cur.children[key] = &WordFilter{}
		}
		cur = cur.children[key]
		cur.prefix = append(cur.prefix, index)
	}
	cur.isEnd = true
}

// from right to left insert the character
func (wf *WordFilter) suffixInsert(index int, word string) {
	cur := wf
	for i := len(word) - 1; i >= 0; i-- {
		key := word[i] - 'a'
		if cur.children[key] == nil {
			cur.children[key] = &WordFilter{}
		}
		cur = cur.children[key]
		cur.suffix = append(cur.suffix, index)
	}
	cur.isEnd = true
}

func (wf *WordFilter) F(pref string, suff string) int {
	suff = reverse(suff)
	p := wf.startWithP(pref)
	s := wf.startWithS(suff)
	if len(p) == 0 || len(s) == 0 {
		return -1
	}
	i, j := len(p)-1, len(s)-1
	for i >= 0 && j >= 0 {
		if p[i] == s[j] {
			return p[i]
		}
		if p[i] > s[j] {
			i--
		} else {
			j--
		}
	}
	return -1
}

// time O(n) space O(1)
func (wf *WordFilter) startWithP(prefix string) []int {
	cur := wf
	for _, ch := range prefix {
		key := ch - 'a'
		if cur.children[key] == nil {
			return []int{}
		}
		cur = cur.children[key]
	}
	return cur.prefix
}

// time O(n) space O(1)
func (wf *WordFilter) startWithS(suffix string) []int {
	cur := wf
	for _, ch := range suffix {
		key := ch - 'a'
		if cur.children[key] == nil {
			return []int{}
		}
		cur = cur.children[key]
	}
	return cur.suffix
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */

// solution 2:
type WordFilter struct {
	children [27]*WordFilter
	isEnd    bool
	index    int
}

func Constructor(words []string) WordFilter {
	root := WordFilter{}
	for index, word := range words {
		word = word + "{"
		for i := 0; i < len(word); i++ {
			cur := &root
			for j := i; j < 2*len(word)-1; j++ {
				k := word[j%len(word)] - 'a'
				if cur.children[k] == nil {
					cur.children[k] = &WordFilter{}
				}
				cur = cur.children[k]
				cur.index = index
				// 相同前缀对应的单词的下标，「小下标」会被「大下标」覆盖
				// 所以，查询匹配出来的就是 largest index
			}
		}
	}
	return root
}

func (root *WordFilter) F(pref string, suff string) int {
	cur := root
	word := suff + "{" + pref
	for _, ch := range word {
		if cur.children[ch-'a'] == nil {
			return -1
		}
		cur = cur.children[ch-'a']
	}
	return cur.index
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(pref,suff);
 */
