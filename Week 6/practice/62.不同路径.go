/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	// dp[i][j] 到第i,j格的路径个数
	// dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	// dp[0][j] = 1
	// dp[i][0] = 1
	dp := make([][]int, m)
	for x := range dp {
		dp[x] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

