package trie

// 212. Word Search II
// 212. 单词搜索 II
// 思路：DFS + 前缀树
// time O(mn*4^10) space O(k*l)
func findWords(board [][]byte, words []string) []string {
	ans := make([]string, 0, len(words))
	r := buildTrie(words)
	m, n := len(board), len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(board, i, j, r, &ans)
		}
	}
	return ans
}

// 208. Implement TrieX (Prefix Tree)
// 208. 实现 TrieX (前缀树)
type TrieX struct {
	Children [26]*TrieX
	IsEnd    bool
	Word     string
}

func buildTrie(words []string) *TrieX {
	root := &TrieX{}
	for _, word := range words {
		root.insert(word)
	}
	return root
}

// time O(n) space O(n)
func (t *TrieX) insert(word string) {
	cur := t
	for _, ch := range word {
		// A link does not exist
		// create a new node and link it with the parent's link
		if cur.Children[ch-'a'] == nil {
			cur.Children[ch-'a'] = &TrieX{}
		}
		// going to the next level
		cur = cur.Children[ch-'a']
	}
	cur.IsEnd = true
	cur.Word = word
}

func dfs(board [][]byte, i, j int, r *TrieX, ans *[]string) {
	// 超出索引范围
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	ch := board[i][j]

	// 走过，不能再走
	if ch == '*' {
		return
	}

	// 与 trie 树节点不匹配
	if r.Children[ch-'a'] == nil {
		return
	}

	// arrived at the end of a word, so find one word
	r = r.Children[ch-'a']
	if r.IsEnd && len(r.Word) != 0 {
		*ans = append(*ans, r.Word)
		r.Word = "" // remove duplication
		// 找到一个单词之后，这里不应该就结束了
		// 后面还有其他单词可以匹配
		// 所以，可以继续 DFS
	}

	// 标记走过
	board[i][j] = '*'

	// 继续往后走
	dfs(board, i-1, j, r, ans) // move up
	dfs(board, i+1, j, r, ans) // move down
	dfs(board, i, j-1, r, ans) // move left
	dfs(board, i, j+1, r, ans) // move right

	// 回溯状态
	board[i][j] = ch
}
