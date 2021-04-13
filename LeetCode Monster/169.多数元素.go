/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
    result := 0
    hashMap := make(map[int]int)
    for _ ,num := range nums {
        hashMap[num]++
        if hashMap[num] > len(nums) / 2{
            result = num
        }
    }
    return result
}
// @lc code=end

