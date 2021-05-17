/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	result, imax, imin := math.MinInt32, 1, 1
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			tmp := imax
			imax = imin
			imin = tmp
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])

		result = max(result, imax)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// @lc code=end

