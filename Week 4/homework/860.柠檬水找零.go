/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	// 维护five，ten的数量
	for _, bill := range bills {

		if bill == 5 { // 5元账单不许找零
			five++
		} else if bill == 10 { // 10元账单，找零5元
			ten++
			five--
		} else if ten > 0 { // 20元账单，找零15元，优先找零10元，然后1张五元
			ten--
			five--
		} else { // 20元账单，找零15元，10元不足，只能找零3张5元
			five = five - 3
		}
		if five < 0 { // 五元没了
			return false
		}
	}
	return true
}

// @lc code=end

