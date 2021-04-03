func climbStairs(n int) int {
	// 记忆化递归
	// 把中间计算过的子问题结果保存下来 map[n] = f(n) // map[int]int
	var tmpMap = make(map[int]int)
	return climbStairsHelper(0,n,tmpMap)

}

func climbStairsHelper(i int, n int, tmpMap map[int]int) int{
	if i == n {
		return 1;
	}
	if i > n {
		return 0;
	}
	// 没有被访问过
	if tmpMap[i] == 0 {
		tmpMap[i] = climbStairsHelper(i + 1, n, tmpMap) + climbStairsHelper(i + 2, n, tmpMap)		
	}

	return tmpMap[i]	
}