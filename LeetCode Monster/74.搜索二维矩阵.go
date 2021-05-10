/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if (matrix == nil || len(matrix) == 0) {
		return false;
	}
	// 当做一维数组
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n-1
	for left <= right {
		mid := (left + right) / 2
		x := matrix[mid/n][mid%n]
		if x == target {
			return true
		} else if x > target {
			right = mid - 1
		} else if x < target {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end

