package main

import "fmt"

func main() {
	words := []string{"appe", "apple"}
	o := Constructor(words)
	r := o.F("a", "le")
	fmt.Println(r)
}

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
