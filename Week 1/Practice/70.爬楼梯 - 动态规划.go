func climbStairs(n int) int {
	//动态规划
	// 转移方程 dp(n) = dp(n - 1) + dp(n - 2)
	// 边界条件 dp(1) = 1 , dp(2) = 2
	if n <= 3 {
		return n
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}
	return dp[n]
	
}