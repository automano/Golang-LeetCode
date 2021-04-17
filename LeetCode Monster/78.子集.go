/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		tmp := make([]int, len(list))
		copy(tmp, list)
		res = append(res, tmp)
		for j := i; j < len(nums); j++ {
			// 做选择
			list = append(list, nums[j])
			//
			dfs(j+1, list)
			// 撤销选择
			list = list[:len(list)-1]
		}
	}

	dfs(0, []int{})
	return res
}

// @lc code=end

