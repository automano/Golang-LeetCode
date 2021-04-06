/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	buckets := make(map[byte]int)
	for i := 0; i < len(secret); i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			if buckets[secret[i]] < 0 {
				cows++
			}
			if buckets[guess[i]] > 0 {
				cows++
			}
			buckets[secret[i]]++
			buckets[guess[i]]--
		}

	}
	result := fmt.Sprintf("%dA%dB", bulls, cows)
	return result
}

// @lc code=end

