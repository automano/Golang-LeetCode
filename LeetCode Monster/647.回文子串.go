/*
 * @lc app=leetcode.cn id=647 lang=golang
 *
 * [647] 回文子串
 */

// @lc code=start
func countSubstrings(s string) int {
	// 转化成二维
	count := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {

			if i == j {
				// 由单个字符组成
				dp[i][j] = true
				count++
			} else if j-i == 1 && s[i] == s[j] {
				// 由 2 个字符组成，且字符要相同
				dp[i][j] = true
				count++
			} else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1] {
				// 由超过 2 个字符组成，首尾字符相同，且剩余子串是一个回文串
				dp[i][j] = true
				count++
			}
		}
	}
	return count
}

// @lc code=end

