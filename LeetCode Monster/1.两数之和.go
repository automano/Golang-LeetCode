/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 2. 哈希映射 时间复杂度是O(n)
	// 如果满足条件， map[target - nums[i]] -> 数组下标
	// 如果不满足， map[nums[i]] = i
	result := make([]int,2)
	tempMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val,ok := tempMap[target - nums[i]]; ok {
			result[0] = val
			result[1] = i
		}
		tempMap[nums[i]] = i
	}
	return result
}	
// @lc code=end