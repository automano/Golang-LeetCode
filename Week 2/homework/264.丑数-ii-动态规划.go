/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start

func nthUglyNumber(n int) int {
	// 2.动态规划
	// dp[n] 第n个丑数
	// dp[1] = 1
	// dp[i] = min(dp[a] * 2 , dp[b] * 3 , dp[c] * 5)
	dp := make([]int, n+1) // 空间复杂度O(n)
	dp[1] = 1
	a, b, c := 1, 1, 1
	for i := 2; i <= n; i++ { // 时间复杂度O(n)
		dp[i] = min(min(dp[a]*2, dp[b]*3), dp[c]*5)
		if dp[i] == dp[a]*2 {
			a++
		}
		if dp[i] == dp[b]*3 {
			b++
		}
		if dp[i] == dp[c]*5 {
			c++
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// @lc code=end

