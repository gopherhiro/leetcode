package trie

// 前缀树
//（只保存小写字符的）「前缀树」是一种特殊的多叉树，
// 它的 TrieNode 中 chidren 是一个大小为 26 的一维数组，
// 分别对应了26个英文字符 'a' ~ 'z'，也就是说形成了一棵 26 叉树。
// 前缀树的结构可以定义为：
// 1. isWord 表示从根节点到当前节点为止，该路径是否形成了一个有效的字符串。
// 2. children 是该节点的所有子节点。

// 208. Implement Trie (Prefix Tree)
// 208. 实现 Trie (前缀树)

type Trie struct {
	children [26]*Trie
	isWord   bool
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
	cur.isWord = true
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
	return cur.isWord
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
