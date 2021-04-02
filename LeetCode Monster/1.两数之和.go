/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 1.暴力解法
	// 两层for循环遍历数组，nums[i] + nums[j] = target? 
	result := make([]int,2) 
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {  //不能重复
			if nums[i] + nums[j] == target {
				result[0] = i
				result[1] = j
			}
		}
	}
	return result;
}
// @lc code=end

