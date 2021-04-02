/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 双指针法 - 缩小搜索空间
	// i 左边界 j 右边界
	// 当左边界高度小于右边界的时候，移动右边界，面积只会减小，此时面积由左边界决定是否会增大，所以要移动左边。
	// 当左边界高度大于右边界的时候，移动左边界，面积只会减小，此时面积由右边界决定时候会增大，所以要移动右边。
	max := 0
	i := 0
	j := len(height) - 1
	for i < j {
		if computeArea(min(height[i],height[j]),(j - i)) > max {
			max = computeArea(min(height[i],height[j]),(j - i))
		}
		// 移动指针
		if height[i] < height[j] {
			i ++
		} else {
			j--
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

