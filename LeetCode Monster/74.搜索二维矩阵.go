/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}
	rowNumber := len(matrix)
	colNumber := len(matrix[0])
	left, right := 0, rowNumber*colNumber-1
	for left < right {
		mid := left + right + 1>>1
		if matrix[mid/rowNumber][mid%colNumber] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return matrix[right/rowNumber][right%colNumber] == target
}

// @lc code=end

