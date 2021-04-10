/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	result := []string{}
	// 定义规则映射表
	hashMap := map[int]string{3: "Fizz", 5: "Buzz"}
	hashMapkeys := []int{3, 5} // 切片保存hashMap key的顺序
	for i := 1; i <= n; i++ {
		// 拼接tempString
		tempString := ""
		for _, key := range hashMapkeys {
			if i%key == 0 {
				tempString += hashMap[key]
			}
		}
		// tempString放入string数组中
		if tempString != "" {
			result = append(result, tempString)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// @lc code=end

