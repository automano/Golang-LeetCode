/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// 动态规划优化
	// 不用dp[]来存状态，因为只跟前两个状态有关
	// 只需要两个变量来保存
	if n < 3{
		return n
	}

	prevprev := 1 
	prev := 2
	cur := 0
	
	for i := 3; i <= n; i++ {
		cur = prevprev + prev
		prevprev = prev
		prev = cur	
	}

	return cur
}
// @lc code=end

