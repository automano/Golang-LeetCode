# 课后总结
## 第5课
### Hash Table
哈希表（Hash table），也叫散列表，是根据关键码值（Key value
）而直接进行访问的数据结构。
它通过把关键码值映射到表中一个位置来访问记录，以加快查找的
速度。
- 这个映射函数叫作散列函数（Hash Function），存放记录的数组叫作哈希表（或散列表）

工程实践
- 用户信息表
- 缓存 LRU Cache
- 键值对存储 Redis

Hash Collisions - 哈希碰撞 （不同的值通过哈希函数映射到相同的值）解决办法：碰撞的元素通过链表存储（最快情况下会退化成链表，时间复杂度退化到O(n))

Golang Code
 
map https://golang.google.cn/ref/spec#Map_types

Go 语言运行时同时使用了多个数据结构组合表示哈希表，其中 runtime.hmap 是最核心的结构体
https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/#33-%E5%93%88%E5%B8%8C%E8%A1%A8

```golang
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap
	nextOverflow *bmap
}
```
Go 语言使用拉链法来解决哈希碰撞的问题实现了哈希表，它的访问、写入和删除等操作都在编译期间转换成了运行时的函数或者方法。哈希在每一个桶中存储键对应哈希的前 8 位，当对哈希进行操作时，这些 tophash 就成为可以帮助哈希快速遍历桶中元素的缓存。

哈希表的每个桶都只能存储 8 个键值对，一旦当前哈希的某个桶超出 8 个，新的键值对就会存储到哈希的溢出桶中。随着键值对数量的增加，溢出桶的数量和哈希的装载因子也会逐渐升高，超过一定范围就会触发扩容，扩容会将桶的数量翻倍，元素再分配的过程也是在调用写操作时增量进行的，不会造成性能的瞬时巨大抖动。

### 面试做题四步
1. clarification
2. possible solution - optimal (time & space)
3. code
4. test cases

## 第6课树，二叉树，二叉搜索树
### 树
Root
Parent Node
Child Node
sub-tree
siblings
left node
right node

如果有环，则为图

### 二叉树
- 只有两个子节点
- 遍历二叉树
  - 前序遍历 Pre-Order : 根-左-右
  - 中序遍历 In-Order: 左-根-右
  - 后序遍历 Post-Order: 左-右-根
- 拥抱递归

### 二叉搜索树
二叉搜索树，也称二叉排序树、有序二叉树（Ordered Binary Tree）、排序二叉树（Sorted Binary Tree），是指一棵空树或者具有下列性质的二叉树：
1. 左子树上所有结点的值均小于它的根结点的值；
2. 右子树上所有结点的值均大于它的根结点的值；
3. 以此类推：左、右子树也分别为二叉查找树。 （这就是 重复性！）
   
中序遍历：升序排列

访问，插入，删除 O(log(n))

## 第6课 堆，二叉堆
### 堆 Heap
可以迅速在一堆数中找到最大值或者最小值的数据结构

大顶堆或小顶堆

常见的堆有：二叉堆和斐波那契堆

golang中package提供了 https://golang.google.cn/pkg/container/heap/

### 二叉堆
通过完全二叉树来实现

二叉堆（大顶）：
- 一颗完全数
- 树中任意节点的值总是》=其子节点的值
  
具体实现：
- 一般通过数组来实现
- 索引 根节点（堆顶元素）是0
  - 索引为i的左孩子的索引是(2*i+1)
  - 索引为i的右孩子的索引是(2*i+2)
  - 索引为i的父节点的索引是(floor(i-1)/2)
- Insert 
  - 新元素一律先插入堆尾部
  - 向上调整顺序，与父亲节点进行比较HeadpifyUp
- Delete Max
  - 讲堆尾的元素替换到顶部
  - 依次从根部向下调整 HedapifyDown

### 图
- Vertices
  - 度：入度和出度
  - 是否连通
- Edges
  - 有向和无向
  - 权重

邻接矩阵和邻接链表

常见的算法
- DFS 要注意不要忘记visited集合
``` python
visited = set() # 和树中的DFS最大区别
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
- BFS
```python
def BFS(graph, start, end):
    queue = [] 
    queue.append([start]) 
    visited = set() # 和树中的BFS的最大区别
    while queue: 
        node = queue.pop() 
        visited.add(node)
        process(node) 
        nodes = generate_related_nodes(node) 
        queue.push(nodes)
```