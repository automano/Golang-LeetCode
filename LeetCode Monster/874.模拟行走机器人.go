/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	// 定义距离
	var dis int
	//初始化机器人的三个参数 位置 x,y 方向 dir
	dir := 0 // 0 -> 北
	disX, disY := 0, 0
	// 输入障碍
	obstaclesMap := make(map[[2]int]bool)
	for _, v := range obstacles {
		if len(v) == 2 {
			obstaclesMap[[2]int{v[0], v[1]}] = true
		}
	}
	//direct* = [4]int{north ,east, south, west}
	directForY := [4]int{1, 0, -1, 0}
	//y轴坐标四个方向的移动如何处理
	directForX := [4]int{0, 1, 0, -1}
	//x轴坐标四个方向的移动如何处理

	// 开始移动
	for len(commands) != 0 {
		// 取出一个指令
		command := commands[0]
		commands = commands[1:]
		// command 为 -1 右转90度
		if command == -1 {
			dis = max(dis, disX*disX+disY*disY)
			dir = (dir + 1) % 4
			continue
		}
		// 转向指令command 为 -2 左转90度
		if command == -2 {
			dis = max(dis, disX*disX+disY*disY)
			dir = (dir + 3) % 4
			continue
		}
		// 移动指令 command >= 0
		for i := 0; i < command; i++ {
			// 一步一步走，判断有没有遇到障碍物
			tempX := disX + directForX[dir]
			tempY := disY + directForY[dir]
			// 遇到障碍物
			if _, block := obstaclesMap[[2]int{tempX, tempY}]; block {
				break
			}
			disX = tempX
			disY = tempY
		}
	}

	return max(dis, disX*disX+disY*disY)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end

