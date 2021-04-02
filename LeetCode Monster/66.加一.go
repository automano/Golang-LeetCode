/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	// 1. 末位+1不需要进位，[4,6] -> [4,7]
	// 2. 末位+1需要进位，中间会停止进位，[4,9,9]->[5,0,0]
	// 3. 末位+1需要进位，不会停止，[9,9] -> [1,0,0]

	// 从末尾开始
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] ++ // +1
		digits[i] %= 10 // 如果进位要变成0，对10取模
		if digits[i] != 0 { // 不再需要进位 
			return digits
		}
	}
	return append([]int{1},digits...) //[1] [0,0]
}
// @lc code=end

