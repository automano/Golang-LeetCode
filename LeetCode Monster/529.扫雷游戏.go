/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	// 初始化
	dirX := [8]int{0, 1, 0, -1, 1, 1, -1, -1}
	dirY := [8]int{1, 0, -1, 0, 1, -1, 1, -1}
	// 深度优先
	var dfs func(x, y int)
	dfs = func(x, y int) {
		count := 0
		// 统计相邻8个位置的地雷数量
		for i := 0; i < 8; i++ {
			tempX := x + dirX[i]
			tempY := y + dirY[i]
			// 超出盘面
			if tempX < 0 || tempX >= len(board) || tempY < 0 || tempY >= len(board[0]) {
				continue
			}
			if board[tempX][tempY] == 'M' {
				count++
			}
		}
		if count > 0 {
			board[x][y] = strconv.Itoa(count)[0]
		} else {
			board[x][y] = 'B'
			// 扫描周围八个点
			for i := 0; i < 8; i++ {
				tempX := x + dirX[i]
				tempY := y + dirY[i]
				// 超出盘面
				if tempX < 0 || tempX >= len(board) || tempY < 0 || tempY >= len(board[0]) || board[tempX][tempY] != 'E' {
					continue
				}
				dfs(tempX, tempY)
			}
		}
	}

	// 开始游戏，点击
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(x, y)
	}
	return board
}

// @lc code=end

