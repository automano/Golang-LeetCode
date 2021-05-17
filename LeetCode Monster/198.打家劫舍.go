/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	// https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/
	// dp[i] = max(dp[i-1],dp[i-2] + nums[i-1])
	// dp[0] = 0
	// dp[1] = nums[0]
	n := len(nums)
	dp := make([]int, n+1)
	// 初始化
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n; i++ {
		// 动态转移方程
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

