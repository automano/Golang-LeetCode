/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	// https://leetcode-cn.com/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/
	// dp[i] = max(dp[i-1],dp[i-2] + nums[i-1])
	// dp[0] = 0
	// dp[1] = nums[0]
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n+1)

	// 两种情况
	// 偷第一间
	nums1 := nums[1:]
	dp[0] = 0
	dp[1] = nums1[0]
	for i := 2; i <= n-1; i++ {
		// 动态转移方程
		dp[i] = max(dp[i-1], dp[i-2]+nums1[i-1])
	}
	p1 := dp[n-1]
	// 不偷第一间
	nums2 := nums[:n]
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= n-1; i++ {
		// 动态转移方程
		dp[i] = max(dp[i-1], dp[i-2]+nums2[i-1])
	}
	p2 := dp[n-1]

	return max(p1, p2)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

