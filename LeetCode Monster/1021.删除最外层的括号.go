/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	result := []byte{}
	level := 0
	for i := 0; i < len(S); i++ {
		if S[i] == ')' {
			level--
		}
		if level >= 1 {
			result = append(result, S[i])
		}
		if S[i] == '(' {
			level++
		}
	}
	return string(result)
}

// @lc code=end

