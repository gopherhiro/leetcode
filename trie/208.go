package trie

// 208. Implement Trie (Prefix Tree)
// 208. 实现 Trie (前缀树)
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

// time O(n) space O(n)
func (t *Trie) Insert(word string) {
	cur := t
	for _, ch := range word {
		key := ch - 'a'
		// A link does not exist
		// create a new node and link it with the parent's link
		if cur.children[key] == nil {
			cur.children[key] = &Trie{}
		}
		// going to the next level
		cur = cur.children[key]
	}
	cur.isEnd = true
}

// time O(n) space O(1)
func (t *Trie) Search(word string) bool {
	cur := t
	for _, ch := range word {
		key := ch - 'a'
		// 节点不匹配，直接 return false
		if cur.children[key] == nil {
			return false
		}
		// 匹配，就去向下一个 level
		cur = cur.children[key]
	}
	// 所有节点都匹配，
	// 就取决于最后一个节点是否是该字符串的末尾
	// Y，匹配
	// N，不匹配
	return cur.isEnd
}

// time O(n) space O(1)
func (t *Trie) StartsWith(prefix string) bool {
	cur := t
	for _, ch := range prefix {
		key := ch - 'a'
		// 节点不匹配，直接 return false
		if cur.children[key] == nil {
			return false
		}
		// 匹配，就去向下一个 level
		cur = cur.children[key]
	}
	// 所有的前缀字符都匹配，表示 trie 中存在 prefix
	return true
}
