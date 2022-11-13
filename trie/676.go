package trie

// 676. Implement Magic Dictionary
// 676. 实现一个魔法字典
type MagicDictionary struct {
	children [26]*MagicDictionary
	isEnd    bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{}
}

func (m *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		m.insert(word)
	}
}

func (m *MagicDictionary) insert(word string) {
	cur := m
	for _, ch := range word {
		if cur.children[ch-'a'] == nil {
			cur.children[ch-'a'] = &MagicDictionary{}
		}
		cur = cur.children[ch-'a']
	}
	cur.isEnd = true
}

func (m *MagicDictionary) Search(searchWord string) bool {
	return dfs(m, searchWord, 0, 1)
}

func dfs(r *MagicDictionary, w string, i int, limit int) bool {
	// base case
	if limit < 0 {
		return false
	}
	if i == len(w) {
		return r.isEnd && limit == 0
	}

	// iterate current node's all children
	ch := w[i] - 'a'
	for c, t := range r.children {
		if t == nil {
			continue
		}
		// c == ch, represent don't need change.
		if c == int(ch) && dfs(t, w, i+1, limit) {
			return true
		}
		// c != ch, represent consume one chance for change
		if c != int(ch) && dfs(t, w, i+1, limit-1) {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
