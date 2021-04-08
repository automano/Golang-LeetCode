/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	//1.哈希表解法
	result := make([]int, 0)
	if nums1 == nil || nums2 == nil {
		return nil
	}
	if len(nums1) <= len(nums2) {
		hashmap1 := make(map[int]int)
		for _, num1 := range nums1 {
			hashmap1[num1]++
		}
		hashmap2 := make(map[int]int)
		for _, nums2 := range nums2 {
			hashmap2[nums2]++
		}
		for key, _ := range hashmap1 {
			if _, ok := hashmap2[key]; ok {
				result = append(result, key)
			}
		}
	} else {
		return intersection(nums2, nums1)
	}
	return result
}

// @lc code=end

