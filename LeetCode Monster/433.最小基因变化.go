/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	// bank []string 放入hashmap中
	bankMap := make(map[string]int)
	for i, s := range bank {
		bankMap[s] = i
	}
	// 判断bank中是否存在目标基因序列
	if _, exist := bankMap[end]; !exist {
		return -1
	}
	// 初始化
	mutationMap := map[byte][]byte{
		'A':[]byte{'C',''}
	}
	step := 0
	queue := []string{start}
	// 如果start在bank中，则将其移除
	if _, exist := bankMap[start]; exist {
		delete(bankMap, start)
	}
	for len(queue) > 0 {
		step++

	}
}

// @lc code=end

