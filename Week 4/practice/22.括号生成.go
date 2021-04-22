/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// 所有可能 -> 回溯算法
	// 回溯算法模板
	// 定义结果集
	result := []string{}
	// 定义回溯函数 dfs(路径，选择列表)
	var dfs func(path []byte, left, right int)
	dfs = func(path []byte, left, right int) {
		// 满足结束条件
		if len(path) == n*2 {
			result = append(result, string(path))
		} else {
			// for 选择 in 选择列表
			if left < n {
				// 做选择
				path = append(path, '(')
				// 调用dfs
				dfs(path, left+1, right)
				// 撤销选择
				path = path[:len(path)-1]
			}
			if right < left {
				// 做选择
				path = append(path, ')')
				// 调用dfs
				dfs(path, left, right+1)
				// 撤销选择
				path = path[:len(path)-1]
			}
		}
	}
	dfs([]byte{}, 0, 0)
	return result
}

// @lc code=end

