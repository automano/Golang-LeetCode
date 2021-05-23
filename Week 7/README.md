# 学习总结
## 第七周
### 复习
#### 初级搜索
 - 朴素搜索
 - 优化方式：不重复（fibonacci）、剪枝（生成括号问题）
 - 搜索方向：
   - DFS: depth first search 深度优先搜索
   - BFS: breadth first search 广度优先搜索
 - 双向搜索
 - 启发式搜索

DFS 代码 - 递归写法
```python
visited = set() 
def dfs(node, visited):
    if node in visited: # terminator
        # already visited 
        return 
    visited.add(node) 
    # process current node here. 
    ...
    for next_node in node.children(): 
        if not next_node in visited: 
            dfs(next_node, visited)
```

DFS 代码 - 非递归写法
```python 
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
BFS 代码
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
```
#### 回溯法

回溯法采用试错的思想，它尝试分步的去解决一个问题。在分步解决问题的过程中，当
它通过尝试发现现有的分步答案不能得到有效的正确的解答的时候，它将取消上一步甚
至是上几步的计算，再通过其它的可能的分步解答再次尝试寻找问题的答案。

回溯法通常用最简单的递归方法来实现，在反复重复上述的步骤后可能出现两种情况:
-   找到一个可能存在的正确的答案
-   在尝试了所有可能的分步方法后宣告该问题没有答案
  
在最坏的情况下，回溯法会导致一次复杂度为指数时间的计算。

#### 双向 BFS
1. https://leetcode-cn.com/problems/word-ladder/
2. https://leetcode-cn.com/problems/minimum-genetic-mutation/

#### 启发式搜索 Heuristic Search (A*)
```python
def AstarSearch(graph, start, end):
    pq = collections.priority_queue() # 优先级 —> 估价函数
    pq.append([start]) 
    visited.add(start)
    while pq: 
        node = pq.pop() # can we add more intelligence here ?
        visited.add(node)
        process(node) 
        nodes = generate_related_nodes(node) 
        unvisited = [node for node in nodes if node not in visited]
        pq.push(unvisited)
```

启发式函数： h(n)，它用来评价哪些结点最有希望的是一个我们要找的结点，h(n) 会返回一个非负实数,也可以认为是从结点n的目标结点路径的估计成本。

启发式函数是一种告知搜索方向的方法。它提供了一种明智的方法来猜测哪个邻居结点会导向一个目标。

1. https://leetcode-cn.com/problems/shortest-path-in-binary-matrix/
2. https://leetcode-cn.com/problems/sliding-puzzle/
3. https://leetcode-cn.com/problems/sudoku-solver/