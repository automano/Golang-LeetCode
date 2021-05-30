/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this // node -> Trie
	for _, ch := range word {
		ch = ch - 'a'
		if node.next[ch] == nil {
			node.next[ch] = &Trie{}
		}
		node = node.next[ch]
	}
	node.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this // node -> Trie
	for _, ch := range word {
		ch = ch - 'a'
		if node.next[ch] == nil {
			return false
		}
		node = node.next[ch]
	}
	return node.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this // node -> Trie
	for _, ch := range prefix {
		ch = ch - 'a'
		if node.next[ch] == nil {
			return false
		}
		node = node.next[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

