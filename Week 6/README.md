# 学习总结
## 第五&六周

### 复习

递归代码模版
``` java
public void recur(int level, int param) { 
    // terminator 
    if (level > MAX_LEVEL) { 
        // process result 
        return; 
    } 
    // process current logic 
    process(level, param); 
    // drill down 
    recur( level: level + 1, newParam); 
    // restore current status 
}
```
分治代码模板

```python
def divide_conquer(problem, param1, param2, ...): 
    # recursion terminator 
    if problem is None: 
        print_result 
        return 
    # prepare data 
    data = prepare_data(problem) 
    subproblems = split_problem(problem, data) 
    # conquer subproblems 
    subresult1 = self.divide_conquer(subproblems[0], p1, ...) 
    subresult2 = self.divide_conquer(subproblems[1], p1, ...) 
    subresult3 = self.divide_conquer(subproblems[2], p1, ...) 
    … 
    # process and generate the final result 
    result = process_result(subresult1, subresult2, subresult3, …) 
    
    # revert the current level states
```

感触
1. 人肉递归低效、很累

2. 找到最近最简方法，将其拆解成可重复解决的问题

3. 数学归纳法思维（抵制人肉递归的诱惑）

本质：寻找重复性 —> 计算机逻辑（只会if，else, loop)

### 动态规划 Dynamic Programming

- wiki 定义 ： https://en.wikipedia.org/wiki/Dynamic_programming

- Simplifying a complicated problem by breaking it down into simpler sub-problems

- Divide & Conquer + Optimal substructure 分治 + 最优子结构

动态规划 和 递归或者分治 没有根本上的区别（关键看有无最优的子结构） 

- 共性：找到重复子问题

- 差异性：最优子结构、中途可以淘汰次优解

动态规划关键点
1. 最优子结构 opt[n] = best_of(opt[n-1], opt[n-2], …)

2. 储存中间状态：opt[i]

3. 递推公式（美其名曰：状态转移方程或者 DP 方程） 

    Fib: opt[i] = opt[n-1] + opt[n-2] 

    二维路径：opt[i,j] = opt[i+1][j] + opt[i][j+1] (且判断a[i,j]是否空地）

实战题目

https://leetcode-cn.com/problems/house-robber/ 

https://leetcode-cn.com/problems/house-robber-ii/description/

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/#/description

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iii/

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-iv/

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/

https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/solution/yi-ge-fang-fa-tuan-mie-6-dao-gu-piao-wen-ti-by-l-3/

https://leetcode-cn.com/problems/perfect-squares/ 

https://leetcode-cn.com/problems/edit-distance/

https://leetcode-cn.com/problems/jump-game/

https://leetcode-cn.com/problems/jump-game-ii/ 

https://leetcode-cn.com/problems/unique-paths/ 

https://leetcode-cn.com/problems/unique-paths-ii/ 

https://leetcode-cn.com/problems/unique-paths-iii/ 

https://leetcode-cn.com/problems/coin-change/

https://leetcode-cn.com/problems/coin-change-2/
    
Homework

https://leetcode-cn.com/problems/longest-valid-parentheses/

https://leetcode-cn.com/problems/minimum-path-sum/

https://leetcode-cn.com/problems/edit-distance/

https://leetcode-cn.com/problems/decode-ways

https://leetcode-cn.com/problems/maximal-square/

https://leetcode-cn.com/problems/max-sum-of-rectangle-no-larger-than-k/

https://leetcode-cn.com/problems/frog-jump/

https://leetcode-cn.com/problems/split-array-largest-sum

https://leetcode-cn.com/problems/student-attendance-record-ii/

https://leetcode-cn.com/problems/task-scheduler/

https://leetcode-cn.com/problems/palindromic-substrings/

https://leetcode-cn.com/problems/minimum-window-substring/

https://leetcode-cn.com/problems/burst-balloons/
