/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}
	result := []int{}
	queue := []int{0} // queue 存储的是数组下标，方便计算

	for i := 1; i < len(nums); i++ {
		for i-queue[0] >= k {
			queue = queue[1:]
		}
		for len(queue) > 0 && nums[i] >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		if i >= k-1 {
			result = append(result, nums[queue[0]])
		}

	}
	return result
}

// @lc code=end

