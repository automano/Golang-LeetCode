/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	ret := []string{}
	// 特殊情况处理
	if len(digits) == 0 {
		return ret
	}
	keyboard := map[byte]string{
		'1': "!@#", '2': "abc", '3': "def",
		'4': "ghi", '5': "jkl", '6': "mno",
		'7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var dfs func(digits string, index int, tempString []byte)
	dfs = func(digits string, index int, tempString []byte) {
		// 满足结束条件
		if index == len(digits) {
			// 添加结果
			ret = append(ret, string(tempString))
		} else {
			digit := digits[index]
			letters := keyboard[digit]
			// for 选择 in 选择列表:
			for i := 0; i < len(letters); i++ {
				// 做选择
				tempString = append(tempString, letters[i])
				// backtrack(路径, 选择列表)
				dfs(digits, index+1, tempString)
				// 撤销选择
				tempString = tempString[:len(tempString)-1]
			}
		}
	}

	dfs(digits, 0, []byte{})
	return ret
}

// @lc code=end

