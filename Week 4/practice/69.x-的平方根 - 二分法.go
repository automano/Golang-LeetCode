/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	// 二分法
	left, right := 0, x
	ans := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			ans = mid
			break
		} else if mid*mid < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}

// @lc code=end

