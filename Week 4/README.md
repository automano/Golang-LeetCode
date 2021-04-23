# 学习总结
## 第四周
### 深度优先遍历&广度优先遍历

遍历 - 搜索
- 访问所有节点
- 且只访问一次
- 访问顺序不限
  - DFS
  - BFS

固化的模板程序 - 记住模板代码，熟练使用

### 深度优先搜索 DFS
```python
# 二叉树
def dfs(node):
    if node in visited:
    #already visited
        return
    visited.add(node)
    # process current node
    # ... # logic here
    dfs(node.Left)
    dfs(node.Right)

# 多叉树
visited = set()
def dfs(node,visited):
    visited.add(node)

    ...
    for next_node in node.Children():
        if not next_node in visited:
            dfs(next_node,visited)

# 非递归写法
# 手动维护一个栈stack
def DFS(self, tree): 
    if tree.root is None: 
        return [] 
        visited, stack = [], [tree.root] 
    while stack: 
        node = stack.pop() 
        visited.add(node) 
        process (node) 
        nodes = generate_related_nodes(node) 
        stack.push(nodes) 
# other processing work 
...
```

### 广度优先遍历 BFS
```python
def BFS(graph, start, end): 
    queue = [] 
    queue.append([start]) 
    visited.add(start) 
    while queue: 
        node = queue.pop() 
        visited.add(node) 
        process(node) 
        nodes = generate_related_nodes(node) 
        queue.push(nodes) 
        # other processing work 
...
```

### 贪心算法 Greedy

贪心算法是一种在每一步选择中都采取在当前状态下最好或最优（即最有利）的选择，从而希望导致结果是全局最好或最优的算法
  
- 与动态规划的区别：贪心算法与动态规划的不同在于它对每个子问题的解决方案都做出选择，不能回退。动态规划则会保存以前的运算结果，并根据以前的结果对当前进行
选择，有回退功能。

- 最优化问题：求图中的最小生成树、求哈夫曼编码等

- 何种情况下用到贪心算法：问题能够分解成子问题来解决，子问题的最优解能递推到最终问题的最优解。这种子问题最优解称为最优子结构。

### 二分查找
#### 前提
   1. 目标函数单调
   2. 存在上下界
   3. 能过够通过索引访问

#### 模板
``` python
left,right = 0,len(array) - 1
while left <= right:
    mid = (left + right) / 2
    if array[mid] == target:
        # find the target!!
        break or return result
        elif array[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
```
