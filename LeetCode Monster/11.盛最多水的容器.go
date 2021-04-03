/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	// 1. 暴力解法 维护一个面积最大值，两层for循环，求两个柱子之间的面积，更新面积最大值，最后输出
	// 时间复杂度 O(n^2) -> 超时
	// 2. 双指针 维护左右两个指针，起始位置左右边界 时间复杂度是O(n)
	// 如果左指针对应的高度小于右指针的，面积受限于左指针的高度，没有增大的可能。这三种情况下就没有必要移动右指针，只需要移动左指针。

	i := 0
	j := len(height) - 1
	max := 0 // 记录面积最大值
	for i < j {
		if (j - i)*min(height[i],height[j]) > max {
			max = (j - i)*min(height[i],height[j]) // 更新面积最大值
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a,b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
// @lc code=end

