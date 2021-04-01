/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	// 普通情况，末尾无进位， 4,6 -> 4,7
	// 末尾有进位 // 进位会中断， 4,9,9 -> 5,0,0
				// 进位不会中断， 9,9 ->  1,0,0	
	// 从末尾开始处理
	for i := len(digits) - 1; i >= 0; i--{
		digits[i] += 1 // +1
		digits[i] %= 10 // 处理进位
		if digits[i] != 0{
			return digits
		}
	}
	// 处理全9的情况
	return append([]int{1},digits...)
}
// @lc code=end

