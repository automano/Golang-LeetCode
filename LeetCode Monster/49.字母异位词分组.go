/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	// 有效异位词升级版 超时
	result := [][]string{}
	isUsed := make(map[string]bool)
	for i := 0; i < len(strs); i++ {
		if !isUsed[strs[i]] {
			tempString := []string{strs[i]}
			for j := i + 1; j < len(strs); j++ {
				if isAnagram(strs[i], strs[j]) {
					isUsed[strs[j]] = true
					tempString = append(tempString, strs[j])
				}
			}
			if len(tempString) > 0 {
				result = append(result, tempString)
			}
		}

	}
	return result
}

// 有效异位词
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

