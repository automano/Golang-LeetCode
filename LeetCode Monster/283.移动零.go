/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// 双指针，快慢指针
	if nums == nil {
		return
	}
	j := 0
	for i := 0; i < len(nums); i++{
		if nums[i] != 0 {
			temp := nums[i]
			nums[i] = nums[j]
			nums[j] = temp
			j++
		}
	}
}
// @lc code=end

