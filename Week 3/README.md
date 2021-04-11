# 学习总结
## 第三周
### 递归
    类似循环 - 通过调用函数体自身进行循环 - 汇编代码很类似

    系统会创建调用栈

#### 递归代码模板
```java
public void recur(int level, int param){
    // recursion terminator 
    if (level > MAX_LEVEL){
        // process result
        return;
    }

    // process current logic
    process(level, param);

    // drill down
    recur(level:level + 1, newParam);

    // restore current status
}
```
#### 思维要点
1. 不要手写递归
2. 找到最近最简的方法，找重复子问题
3. 数学归纳法的思维

#### 技巧
- mutual exclusive
- complete exhaustive

### 分治， 回溯

#### 代码模板
``` python
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
    ...
    # process and generate the final result 
    result = process_result(subresult1, subresult2, subresult3, …) 
 
 # revert the current level states
 ```