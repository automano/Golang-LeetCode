/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num < 2 {
		return true
	}
	x := num / 2
	for x*x > num {
		x = (x + num/x) / 2
	}
	return x*x == num
}

// @lc code=end

