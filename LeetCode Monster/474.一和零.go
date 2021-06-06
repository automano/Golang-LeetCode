/*
 * @lc app=leetcode.cn id=474 lang=golang
 *
 * [474] 一和零
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	len := len(strs)
	cnt := make([][2]int, len)
	// cnt 记录每个strs元素中0，1的个数
	for i, str := range strs {
		zeros, ones := 0, 0
		for ch := range str {
			if ch == '0' {
				zeros++
			} else {
				ones++
			}
		}
		cnt[i] = [2]int{zeros, ones}
	}
	// 定义dp数据切片
	dp := make([][][]int, len+1)
	for k := range dp {
		dp[k] = make([][]int, m+1)
		for i := range dp[k] {
			dp[k][i] = make([]int, n+1)
		}
	}
	// 状态转移
	for k := 1; k <= len; k++ {
		zeros, ones := cnt[k-1][0], cnt[k-1][1]
		for i := 0; i <= m; i++ {
			for j := 0; j <= n; j++ {
				if i >= zeros && j >= ones {
					dp[k][i][j] = max(dp[k-1][i][j], dp[k-1][i-zeros][j-ones]+1)
				} else {
					dp[k][i][j] = max(dp[k-1][i][j], 0)
				}
			}
		}
	}

	return dp[len][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

