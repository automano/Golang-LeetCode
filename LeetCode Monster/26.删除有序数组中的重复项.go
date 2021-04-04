/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if nums == nil {
		return 0
	}
	i := 0
	j := 1
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			nums[i] = nums[j]
			i = j
		}
	}
	return i + 1
}
// @lc code=end

