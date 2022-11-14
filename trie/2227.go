package trie

import "strings"

// 2227. Encrypt and Decrypt Strings
// 2227. 加密解密字符串
// 思路：逆向思维
type Encrypter struct {
	children [26]*Encrypter
	isEnd    bool
	count    int             // counting the number of dictionary's word same encrypt code
	m        map[byte]string // keys[i]:values[i]
}

func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
	e := Encrypter{}
	// keys[i]:values[i]
	e.m = make(map[byte]string, len(keys))
	for i := 0; i < len(keys); i++ {
		e.m[keys[i]] = values[i]
	}
	// encrypt the words of the dictionary
	// and store them in a trie
	// with a count of each words.
	for _, word := range dictionary {
		w := e.Encrypt(word)
		if len(w) == 0 {
			continue
		}
		e.insert(w)
	}
	return e
}

func (e *Encrypter) insert(w string) {
	cur := e
	for _, ch := range w {
		if cur.children[ch-'a'] == nil {
			cur.children[ch-'a'] = &Encrypter{}
		}
		cur = cur.children[ch-'a']
	}
	cur.isEnd = true
	cur.count++
}

func (e *Encrypter) search(w string) bool {
	cur := e
	for _, ch := range w {
		if cur.children[ch-'a'] == nil {
			return false
		}
		cur = cur.children[ch-'a']
	}
	return cur.isEnd
}

// time O(n) space O(n)
func (e *Encrypter) Encrypt(w string) string {
	ans := strings.Builder{}
	for _, ch := range w {
		val, ok := e.m[byte(ch)]
		if !ok {
			// in case a character of the string is not present in keys,
			// the encryption process cannot be carried out,
			// and an empty string "" is returned.
			return ""
		}
		ans.WriteString(val)
	}
	return ans.String()
}

// time O(n) space O(n)
func (e *Encrypter) Decrypt(w string) int {
	// traversing over the trie
	// with decrypted word to get the count.
	cur := e
	for _, ch := range w {
		if cur.children[ch-'a'] == nil {
			return 0
		}
		cur = cur.children[ch-'a']
	}
	return cur.count
}

/**
 * Your Encrypter object will be instantiated and called as such:
 * obj := Constructor(keys, values, dictionary);
 * param_1 := obj.Encrypt(word1);
 * param_2 := obj.Decrypt(word2);
 */
