/*
 * @lc app=leetcode.cn id=221 lang=golang
 *
 * [221] 最大正方形
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	// dp[i][j] - i,j为右下角的正方形最大值
	// 初始化 
	// 状态转移 dp[i][j] = min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) + 1
	reuslt := 0
	dp = matrix
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp
		}
	}

}
// @lc code=end

