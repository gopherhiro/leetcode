package trie

// 211. 添加与搜索单词 - 数据结构设计
// 211. Design Add and Search Words Data Structure
type WordDictionary struct {
	Children [26]*WordDictionary
	IsEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

// time O(n) space O(n)
func (wd *WordDictionary) AddWord(word string) {
	cur := wd
	for _, ch := range word {
		ch -= 'a'
		if cur.Children[ch] == nil {
			cur.Children[ch] = &WordDictionary{}
		}
		cur = cur.Children[ch]
	}
	cur.IsEnd = true
}

// recursion solution
// time O(n) space o(n)
func (wd *WordDictionary) Search(word string) bool {
	return dfs(wd, word, 0)
}

func dfs(r *WordDictionary, w string, i int) bool {
	// base case
	if i == len(w) {
		return r.IsEnd
	}
	// w[i] is  '.' , we need to iterate the children
	// to find what is not empty node to dfs
	if w[i] == '.' {
		for _, t := range r.Children {
			if t == nil {
				continue
			}
			if dfs(t, w, i+1) {
				return true
			}
		}
		return false
	}
	// w[i] is not '.' , just compare current node
	// and go to the children do the same things
	ch := w[i] - 'a'
	if r.Children[ch] == nil {
		return false
	}
	r = r.Children[ch]
	return dfs(r, w, i+1)
}
