/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 暴力解法
	// 两层for 循环 O(n^2) 超时
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			// 面积 = height中较小的一个 * 边长（j - i)
			if computeArea(min(height[i],height[j]),j - i) > max {
				// 更新面积最大值 max
				max = computeArea(min(height[i],height[j]),j - i)
			}
		}
	}
	return max
}

// 求最小值
func min(a,b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}

// 计算面积
func computeArea(a,b int) int {
	return a * b
}
// @lc code=end

