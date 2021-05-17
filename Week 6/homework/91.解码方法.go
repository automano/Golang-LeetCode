/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	// 动态规划 dp[i] - 第i个字符结尾的字符串所表示的组合个数
	// dp[0] = 0
	// dp[i] = dp[i - 1]  s[i] 独立一个字母

	n := len(s)
	s = " " + s
	dp := make([]int,3)
	dp[0] = 1
	for i := 1; i <= n; i++{
		dp[i%3]=0
		a, b := s[i] - '0', (s[i - 1] - '0') * 10 + (s[i] - '0')
		if a >=1&&a<=9 {
			dp[i%3] = dp[(i-1)%3]
		}
		if b >=10 &&b <=26 {
			dp[i%3] += dp[(i-2)%3]
		}
	}

	return dp[n%3]
}
// @lc code=end

