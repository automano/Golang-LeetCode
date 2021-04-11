/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	// 3. 数学解法 - 数字根
	// 数字根 % 9 == num % 9
	return (num-1)%9 + 1

}

// @lc code=end

