/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ret := []string{}
	if n == 0 {
		return ret
	}
	dfs("", n, n, &ret)
	return ret
}

// dfs
func dfs(cur string, left int, right int, ret *[]string) {
	if left == 0 && right == 0 {
		*ret = append(*ret, cur)
		return
	}

	// 剪枝
	if left > right {
		return
	}
	if left > 0 {
		dfs(cur+"(", left-1, right, ret)
	}

	if right > 0 {
		dfs(cur+")", left, right-1, ret)
	}
}

// @lc code=end

