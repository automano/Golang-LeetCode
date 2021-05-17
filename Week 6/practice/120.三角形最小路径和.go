/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	// dp[i][j] 从i，j到底边的最小路径和
	// dp[i][j] = min(dp[i+1][j],dp[i+1][j+1]) + triangle[i][j]
	// bottom-up
	dp := triangle
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + dp[i][j]
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

// @lc code=end

