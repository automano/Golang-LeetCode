/*
 * @lc app=leetcode.cn id=231 lang=golang
 *
 * [231] 2的幂
 */

// @lc code=start
func isPowerOfTwo(n int) bool {
	// 一个数 n 是 2 的幂，当且仅当 n 是正整数，并且 n 的二进制表示中仅包含 1 个 1。
	// n & (n - 1) -> 直接将 n 二进制表示的最低位 1 移除
	return n > 0 && n&(n-1) == 0
}

// @lc code=end

