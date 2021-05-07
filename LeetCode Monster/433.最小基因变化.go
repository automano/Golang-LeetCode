/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
func minMutation(start string, end string, bank []string) int {
	base := [4]byte{'A', 'C', 'G', 'T'}
	hashGenBank := make(map[string]bool, len(bank))
	for _, gen := range bank {
		hashGenBank[gen] = true
	}
	queue := []string{start}
	minLevel := 0
	// BFS
	for len(queue) != 0 {
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			gen := queue[0]
			queue = queue[1:]
			if gen == end {
				return minLevel
			}

			for g := 0; g < len(gen); g++ {
				for j := 0; j < 4; j++ {
					newGen := gen[:g] + string(base[j]) + gen[g+1:]
					if hashGenBank[newGen] {
						queue = append(queue, newGen)
						hashGenBank[newGen] = false
					}
				}
			}
		}
		minLevel++
	}
	return -1
}

// @lc code=end

