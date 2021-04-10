/*
 * @lc app=leetcode.cn id=263 lang=golang
 *
 * [263] 丑数
 */

// @lc code=start
func isUgly(n int) bool {
	// n = 2^a + 3^b + 5^c
	// n / 2
	for n <= 0 {
		return false
	}

	for n%2 == 0 {
		n = n / 2
	}

	for n%3 == 0 {
		n = n / 3
	}

	for n%5 == 0 {
		n = n / 5
	}

	if n == 1 {
		return true
	}

	return false
}

// @lc code=end

