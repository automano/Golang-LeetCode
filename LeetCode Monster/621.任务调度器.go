/*
 * @lc app=leetcode.cn id=621 lang=golang
 *
 * [621] 任务调度器
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	max := 0
	maxCount := 0
	s := [26]int{}
	for _, v := range tasks {
		s[v-'A']++
		if s[v-'A'] == max {
			maxCount++
		} else if s[v-'A'] > max {
			max = s[v-'A']
			maxCount = 1
		}
	}
	maxTime := (max-1)*(n+1) + maxCount
	if maxTime < len(tasks) {
		maxTime = len(tasks)
	}
	return maxTime
}

// @lc code=end

