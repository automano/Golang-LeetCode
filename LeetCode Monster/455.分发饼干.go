/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	count := 0
	// 特殊情况
	if g == nil || s == nil {
		return count
	}
	// 正常情况
	sort.Ints(g) // 小孩胃口
	sort.Ints(s) // 饼干大小
	// 先满足胃口小的孩子
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}

// @lc code=end

