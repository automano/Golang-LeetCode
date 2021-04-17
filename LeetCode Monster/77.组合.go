/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	ret := [][]int{}
	// 特殊情况
	if k <= 0 || n < k {
		return ret
	}
	var dfs func(index int, path []int)
	dfs = func(index int, path []int) {
		if len(path) == k {
			temp := make([]int, len(path))
			copy(temp, path)
			ret = append(ret, temp)
		} else {
			for i := index; i <= n; i++ {
				path = append(path, i)
				dfs(i+1, path)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(1, []int{})
	return ret
}

// @lc code=end

