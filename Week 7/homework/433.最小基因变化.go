/*
 * @lc app=leetcode.cn id=433 lang=golang
 *
 * [433] 最小基因变化
 */

// @lc code=start
var mode = []byte{'A', 'C', 'G', 'T'}

func minMutation(start string, end string, bank []string) int {
	st := map[string]struct{}{}
	et := map[string]struct{}{}
	dict := map[string]struct{}{}

	// 合法序列
	for _, b := range bank {
		dict[b] = struct{}{}
	}
	// 目标序列不合法
	if _, ok := dict[end]; !ok {
		return -1
	}
	st[start] = struct{}{}
	et[end] = struct{}{}

	return bfs(st, et, dict, 1)
}

func bfs(st, et, dict map[string]struct{}, step int) int {
	if len(st) == 0 {
		return -1
	}
	// 从数量少的一端搜索
	if len(st) > len(et) {
		return bfs(et, st, dict, step)
	}
	next := map[string]struct{}{}
	// 遍历起始序列 处理当前层(广度优先遍历)
	for s, _ := range st {
		tmp := []byte(s)
		// 遍历起始序列每个字符
		for i := 0; i < 8; i++ {
			// 遍历所有可能的变化
			for j := 0; j < 4; j++ {
				// 相同则跳过
				if tmp[i] == mode[j] {
					continue
				}
				// 变化字符
				tmp[i] = mode[j]
				// 判断变化后序列是否为目的序列
				if _, ok := et[string(tmp)]; ok {
					return step
				}
				// 判断变化后序列是否合法
				if _, ok := dict[string(tmp)]; ok {
					// 将合法序列添加到下次遍历集合
					next[string(tmp)] = struct{}{}
				}
				// 清除变化
				tmp[i] = s[i]
			}
		}
	}
	// 探索下一层
	return bfs(next, et, dict, step+1)
}

// @lc code=end

