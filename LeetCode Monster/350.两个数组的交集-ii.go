/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	//2. 先排序（已经排序玩的） 双指针解法
	sort.Ints(nums1) // 1,1,2,2
	sort.Ints(nums2) // 2,2
	i, j := 0, 0
	result := make([]int, 0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		} else {
			result = append(result, nums1[i])
			i++
			j++
		}
	}
	return result
}

// @lc code=end

