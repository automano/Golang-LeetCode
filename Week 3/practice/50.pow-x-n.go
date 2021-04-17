/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	var ret float64
	if n >= 0 {
		ret = pow(x,n)
	} else{
		ret = 1 / pow(x, -n)
	}
	return ret
}

func pow (x float64,n int) float64 {
	if n == 0 {
		return 1.0
	}

	temp := pow(x,n/2)

	if n%2 == 0 {
		return temp*temp
	} else {
		return temp*temp*x
	}
}
// @lc code=end

