/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	//动态规划
	// 转移方程 dp(n) = dp(n - 1) + dp(n - 2)
	// 边界条件 dp(1) = 1 , dp(2) = 2
	// 没有必要使用dp数据，当前状态只跟前两个状态有关，保存为两个变量
	if n <= 3 {
		return n
	}
	var prevprev = 1 // dp(1)
	var prev = 2 // dp(2)
	var cur = 0 // dp[n]
	for i := 3; i <= n; i++ {
		cur = prevprev + prev
		prevprev = prev
		prev = cur
	}
	return cur	
}
// @lc code=end

