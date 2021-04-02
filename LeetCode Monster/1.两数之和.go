/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	// 使用一个哈希映射 map[int] int 来存 key->nums[i], value->i
	// 把循环缩减到1次
	result := make([]int,2)
	tempMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		// 判断tempMap中是否存在 target - nums[i] 的key
		// 如果存在, i,tempMap[target - nums[i]] ->result
		// 如果不存在，tempMap[num[i]] = i
		if val, ok := tempMap[target - nums[i]]; ok {
			result[0] = i
			result[1] = val
		}
		tempMap[nums[i]] = i
	}
	return result
}
// @lc code=end

