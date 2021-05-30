# 学习总结

## 第八周 - 第十四课 - 字典树和并查集

### 字典树 Trie

#### 基本结构

字典树，即Trie树，又称为单词查找树或者键树是一种树形结构。典型应用适用于统计和排序的大量的字符串，所以经常被搜索引擎用于文本词频统计。

它的优点是：最大限度地减少无谓的字符串比较，查询效率比哈希表高。


#### 基本性质

1. 节点本身不存完成单词
2. 从根节点到某一结点，路径上经过的字符连接起来，为该结点对应的字符串
3. 每个结点的所有子节点路径代表的字符都不相同
   
   节点可以用来存储额外的信息


#### 核心思想

Trie树的核心思想是空间换时间
利用字符串的公共前缀来降低查询时间的开销以达到提高效率的目的


#### 模板代码

```python
# Python class Trie(object):
def __init__(self):
    self.root = {} 	
    self.end_of_word = "#"  	
    def insert(self, word): 		
        node = self.root 		
        for char in word: 			
            node = node.setdefault(char, {}) 		
            node[self.end_of_word] = self.end_of_word  	
            def search(self, word): 		
                node = self.root 		
                for char in word: 			
                    if char not in node: 				
                        return False 			
                        node = node[char] 		
                        return self.end_of_word in node  	
                        def startsWith(self, prefix): 		
                            node = self.root 		
                            for char in prefix: 			
                                if char not in node: 				return False 			
                                node = node[char] 		
                                return True
```


#### 实战题目

1. https://leetcode-cn.com/problems/implement-trie-prefix-tree/#/description
2. https://leetcode-cn.com/problems/word-search-ii/
3. Search suggestion - system design



### 并查集 - Disjoint Set

#### 适用场景
- 组团，配对问题
- Group？or not?
  

#### 基本操作

- makeSet(s) : 建立一个新的并查集，其中包含s个单元素集合
- unionSet(x,y): 把元素x和元素y所在的集合合并，要求x和y所在的集合不相交，如果相交则不合并
- find(x):  找到元素x所在的集合的代表，该操作用于判断两个元素是否位于同一个集合，只要找到他们各自的代表比较一下就可以。


### 基本步鄹

- 初始化： 自身指向自身
- 查询、合并
- 路径压缩


#### 模板代码

```python
# Python 
def init(p):
    # for i = 0 .. n: p[i] = i;
    p = [i for i in range(n)]
    def union(self, p, i, j):
        p1 = self.parent(p, i)
        p2 = self.parent(p, j)
        p[p1] = p2
        def parent(self, p, i):
            root = i
            while p[root] != root:
                root = p[root]
                while p[i] != i:
                    # 路径压缩 ?
                    x = i;
                    i = p[i];
                    p[x] = root
                    return root
```


#### 实战题目

- https://leetcode-cn.com/problems/friend-circles
- https://leetcode-cn.com/problems/number-of-islands/
- https://leetcode-cn.com/problems/surrounded-regions/



## 第八周 - 第十五课 - 高级树、AVL 树和红黑树

### 二叉搜索树 Binary Search Tree

二叉搜索树，也称二叉搜索树、有序二叉树（Ordered Binary Tree）、排
序二叉树（Sorted Binary Tree），是指一棵空树或者具有下列性质的二叉
树：
1. 左子树上所有结点的值均小于它的根结点的值；
2. 右子树上所有结点的值均大于它的根结点的值；
3. 以此类推：左、右子树也分别为二叉查找树。 （这就是 重复性！）
   

中序遍历：升序排列

保证性能的关键：

1. 保证二维维度！ —> 左右子树结点平衡（recursively）
2. Balanced
3. https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree

### AVL 树

1. 发明者 G. M. Adelson-Velsky和 Evgenii Landis
2. Balance Factor（平衡因子）是它的左子树的高度减去它的右子树的高度（有时相反）。记录左右子树高度，balance factor = {-1, 0, 1} 
3. 通过旋转操作来进行平衡（四种）
   
    https://zhuanlan.zhihu.com/p/63272157

    https://en.wikipedia.org/wiki/Tree_rotation#/media/File:Tree_Rebalancing.gif

   - 左旋
   - 右旋
   - 左右旋
   - 右左旋
   
4. https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree

不足：结点需要存储额外信息、且调整次数频繁

### Red-black Tree
红黑树是一种近似平衡的二叉搜索树（Binary Search Tree），它能够确保任何一个结点的左右子树的高度差小于两倍。具体来说，红黑树是满足如下条件的二叉搜索树：
- 每个结点要么是红色，要么是黑色
- 根结点是黑色
- 每个叶结点（NIL结点，空结点）是黑色的
- 不能有相邻接的两个红色结点
- 从任一结点到其每个叶子的所有路径都包含相同数目的黑色结点。

#### 关键性质

从根到叶子的最长的可能路径不多于最短的可能路径的两倍长。

#### 对比
- AVL trees providefaster lookups than Red Black Trees because they are more strictly balanced.
- Red Black Trees providefaster insertion and removal operations than AVL trees as fewer rotations are done due to relatively relaxed balancing.
- AVL trees storebalance factors or heightswith each node, thus requires storage for an integer per node whereas Red Black Tree requires only 1 bit of information per node.
- Red Black Trees are used in most of the language libraries 
likemap, multimap, multisetin C++whereas AVL trees are used indatabaseswhere faster retrievals are required


## 第八周 - 第十六课 -  位运算

#### 为什么需要位运算

- 机器里的数字表示方式和存储格式就是 二进制
- 十进制 <—> 二进制 : 如何转换？
https://zh.wikihow.com/%E4%BB%8E%E5%8D%81%E8%BF%9B%E5%88%B6%E8%BD%AC%E6%8D%A2%E4%B8%BA%E4%BA%8C%E8%BF%9B%E5%88%B6

#### 位运算符

- 左移 << 0011 -> 0110
- 右移 >> 0110 -> 0011
- 按位或 | 
- 按位与 &
- 按位取反 ~
- 按位异或 ^

#### XOR - 异或

异或： 相同为0，不同为1. 也可用不进位加法来理解。

特点：
- x^0 = x
- x^1s = x^~0= ~x
- x^~x = 1s
- x^x = 0
- c = a^b -> a^c = b -> b^c = a
- a^b^c = a^(b^c) = (a^b)^c

#### 指定位置的位运算
1. 将 x 最右边的 n 位清零：x & (~0 << n)
2. 获取 x 的第 n 位值（0 或者 1）： (x >> n) & 1
3. 获取 x 的第 n 位的幂值：x & (1 << n)
4. 仅将第 n 位置为 1：x | (1 << n)
5. 仅将第 n 位置为 0：x & (~ (1 << n))
6. 将 x 最高位至第 n 位（含）清零：x & ((1 << n) - 1)

#### 实战位运算要点
- 判断奇偶：
  - x % 2 == 1 —> (x & 1) == 1
  - x % 2 == 0 —> (x & 1) == 0

- x >> 1 —> x / 2
  - x = x / 2; —> x = x >> 1;
  - mid = (left + right) / 2; —> mid = (left + right) >> 1;

- X = X & (X-1) 清零最低位的 1
- X & -X => 得到最低位的 1 
- X & ~X => 0

#### 实战题目
- https://leetcode-cn.com/problems/number-of-1-bits/
- https://leetcode-cn.com/problems/power-of-two/
- https://leetcode-cn.com/problems/reverse-bits/
- https://leetcode-cn.com/problems/n-queens/description/
- https://leetcode-cn.com/problems/n-queens-ii/description/