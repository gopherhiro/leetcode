package trie

import "strings"

type Trie struct {
	children [26]*Trie
	word     string
	isEnd    bool
}

// 648. Replace Words
// 648. 单词替换
// 思路：trie
// time O(n) space O(n)
func replaceWords(dictionary []string, sentence string) string {
	// put all the roots in a trie
	root := &Trie{}
	for _, word := range dictionary {
		cur := root
		for _, ch := range word {
			if cur.children[ch-'a'] == nil {
				cur.children[ch-'a'] = &Trie{}
			}
			cur = cur.children[ch-'a']
		}
		cur.word = word
		cur.isEnd = true
	}

	// iterate trie and match the prefix of a word
	answer := &strings.Builder{}
	s := strings.Split(sentence, " ")
	for _, word := range s {
		if len(answer.String()) > 0 {
			answer.WriteString(" ")
		}
		cur := root
		for _, ch := range word {
			// 不匹配，表示当前单词不用替换，直接拼接即可。
			if cur.children[ch-'a'] == nil {
				break
			}
			// 匹配一个词根，则拼接词跟即可。
			if cur.isEnd {
				break
			}
			cur = cur.children[ch-'a']
		}
		if cur.isEnd {
			answer.WriteString(cur.word)
		} else {
			answer.WriteString(word)
		}
	}
	return answer.String()
}

// 648. Replace Words
// 648. 单词替换
// 思路：hashtable
// time O(len(sentence)) space O(n)
func replaceWordsX(dictionary []string, sentence string) string {
	seen := make(map[string]bool, 0)
	for _, root := range dictionary {
		seen[root] = true
	}

	answer := &strings.Builder{}
	s := strings.Split(sentence, " ")
	for _, word := range s {
		// 构造的字符串是以空格隔开的
		if len(answer.String()) > 0 {
			answer.WriteString(" ")
		}
		prefix := ""
		for i := 1; i <= len(word); i++ {
			// 遇到匹配的前缀后，就 break 停止该单词的后续可能前缀的匹配
			// 使得可以获得长度最短的词根去替换掉 sentence 中的单词
			prefix = word[:i]
			if seen[prefix] {
				break
			}
		}
		answer.WriteString(prefix)
	}
	return answer.String()
}
