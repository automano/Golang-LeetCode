/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	count := 0
	if grid == nil {
		return count
	}
	// 边界
	row, col := len(grid), len(grid[0])
	// 找到该点附近的所有陆地
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 终止条件
		if i < 0 || j < 0 || i >= row || j >= col || grid[i][j] == '0' {
			return
		}
		// 处理当前层逻辑
		grid[i][j] = '0'
		// dill down
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// @lc code=end

