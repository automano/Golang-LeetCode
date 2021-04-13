/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N 皇后
 */
var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	// 行数到n意味着找到所有N皇后
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		// 该列已有queen占据
		if columns[i] {
			continue
		}
		// 主对角线已经被占据
		if diagonals1[row-i] {
			continue
		}
		// 副对角线已经被占据
		if diagonals2[row+i] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[row-i], diagonals2[row+i] = true, true
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, row-i)
		delete(diagonals2, row+i)
	}
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

