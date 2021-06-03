/*
 * @lc app=leetcode.cn id=523 lang=golang
 *
 * [523] 连续的子数组和
 */

// @lc code=start
func checkSubarraySum(nums []int, k int) bool {
	// 维护前缀和
	sum := make([]int, len(nums)+1)
	sum[0] = nums[0]
	for i := 1; i <= len(nums); i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	//
	hashMap := make(map[int]bool, len(nums))
	for i := 2; i <= len(nums); i++ {
		hashMap[sum[i-2]%k] = true
		if hashMap[sum[i]%k] {
			return true
		}
	}
	return false
}

// @lc code=end

