/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int)  {
	// j 记录非零的个数
	// 循环的时候遇到0就跳过，非零的就和当前的i进行交换
	if nums == nil{
		return
	}
	j := 0
	for i := 0; i < len(nums); i++{
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	// 把j个0写到数组末尾
	for i := j; i < len(nums); i++ {
		nums[i] = 0
	}
}
// @lc code=end

