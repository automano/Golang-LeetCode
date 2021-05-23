func climbStairs(n int) int {
	// 1. 定义子问题
	// 爬第n阶的方法取决于第n-1阶和第n-2阶
	// 2. 写出子问题的递推关系
	// dp[i] = dp[i-1] + dp[i-2]
	// 3. 确定 DP 数组的计算顺序
	// 从左往右计算
	// 4. 空间优化（可选）
	// 仅保存最近两个台阶的方法
	if n < 3 {
		return n
	}
	prepre := 1
	pre := 2
	cur := 0
	for i := 3; i <= n; i++ {
		cur = prepre + pre
		prepre = pre
		pre = cur
	}
	return cur
}