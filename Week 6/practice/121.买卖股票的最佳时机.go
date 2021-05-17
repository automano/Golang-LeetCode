/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	result := 0
	if len(prices) == 0 {
		return result
	}
	minPrice := math.MaxInt64

	// 只遍历一遍，找到最大值，最小值，作差
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		}
		tempPrice := price - minPrice
		if tempPrice > result {
			result = tempPrice
		}
	}
	return result

}

// @lc code=end

