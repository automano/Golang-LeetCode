/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
func solveNQueens(n int) [][]string {
	// 回溯算法
	// result = []
	// def backtrack(路径, 选择列表):
	//     if 满足结束条件:
	//         result.add(路径)
	//         return
	//     for 选择 in 选择列表:
	//         做选择
	//         backtrack(路径, 选择列表)
	//         撤销选择
	// 结果
	result := [][]string{}
	// queens所在的行
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	// queens所在的列
	columns := map[int]bool{}
	// queens所在的对角线
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	var dfs func(row int)
	dfs = func(row int) {
		if n == row {
			result = append(result, generateBoard(queens, n))
		}
		for i := 0; i < n; i++ {
			// 做选择
			if columns[i] || diagonals1[row-i] || diagonals2[row+i] {
				continue
			}
			queens[i] = 0
			columns[i], diagonals1[row-i], diagonals2[row+i] = true, true, true
			dfs(row + 1)
			// 撤销选择
			queens[i] = -1
			delete(columns, i)
			delete(diagonals1, row-i)
			delete(diagonals2, row+i)
		}
	}
	dfs(0)
	return result
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
}

// @lc code=end

