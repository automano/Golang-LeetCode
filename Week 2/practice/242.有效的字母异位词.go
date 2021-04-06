/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// 哈希表解法
	if len(s) != len(t) {
		return false
	}
	// 初始化哈希表维护字符值出现的次数
	tempMap := make(map[rune]int) // 考虑unicode,在golang中使用rune来遍历字符串

	for _, ch := range s {
		tempMap[ch]++
	}
	for _, ch := range t {
		tempMap[ch]--
	}
	for _, value := range tempMap {
		if value != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

