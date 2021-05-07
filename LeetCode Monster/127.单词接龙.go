/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	hashWordList := make(map[string]bool, len(wordList))
	for _, s := range wordList {
		hashWordList[s] = true
	}
	queue := []string{beginWord}
	minLevel := 1
	// BFS
	for len(queue) != 0 {
		queueSize := len(queue)
		// 取出同层所有元素
		for i := 0; i < queueSize; i++ {
			// 从队首取出一个元素
			word := queue[0]
			queue = queue[1:]
			if word == endWord {
				return minLevel
			}

			// 构造新单词，判断是否存在于wordlist
			for w := 0; w < len(word); w++ {
				for j := 'a'; j <= 'z'; j++ {
					newWord := word[:w] + string(j) + word[w+1:]
					if hashWordList[newWord] {
						queue = append(queue, newWord)
						hashWordList[newWord] = false
					}
				}
			}
		}
		minLevel++
	}
	return 0
}

// @lc code=end

