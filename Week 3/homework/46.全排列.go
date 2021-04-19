/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	// 回溯算法
	// 定义结果
	result := [][]int{}
	visited := make([]bool, len(nums))
	// 定义回溯函数 dfs(路径，选择列表)
	var dfs func(path []int, index int)
	dfs = func(path []int, index int) {
		// 满足结束条件
		if index == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			result = append(result, temp)
		} else {
			// for 选择 in 选择列表
			for i := 0; i < len(nums); i++ {
				if visited[i] {
					continue
				}
				// 做选择
				path = append(path, nums[i])
				visited[i] = true
				dfs(path, index+1)
				// 撤销选择
				path = path[:len(path)-1]
				visited[i] = false
			}
		}
	}
	dfs([]int{}, 0)
	return result
}

// @lc code=end

