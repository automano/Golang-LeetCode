# 学习总结

## 第九周 - 第十七课 - 布隆过滤器和LRU Cache

### 布隆过滤器

#### 什么情况下需要布隆过滤器

- 网页爬虫对URL的去重，避免爬取相同的URL地址；
- 反垃圾邮件，从数十亿个垃圾邮件列表中判断某邮箱是否垃圾邮箱（同理，垃圾短信）
- 缓存击穿，将已存在的缓存放到布隆中，当黑客访问不存在的缓存时迅速返回避免缓存及DB挂掉

一句话：  如何判断一个元素是否存在一个集合中？

#### 常规思路
- 数组
- 链表
- 树、平衡二叉树、Trie
- Map (红黑树)
- 哈希表

但假如数据量特别大，500万或者1亿条数据。这个时候常规的数据结构的问题就凸显出来了。数组、链表、树等数据结构会存储元素的内容，一旦数据量过大，消耗的内存也会呈现线性增长，最终达到瓶颈。有的同学可能会问，哈希表不是效率很高吗？查询效率可以达到O(1)。但是哈希表需要消耗的内存依然很高。

#### 布隆过滤器介绍

- 巴顿.布隆于一九七零年提出
- 一个很长的二进制向量 （位数组）
- 一系列随机函数 (哈希)
- 空间效率和查询效率高
- 有一定的误判率（哈希表是精确匹配）

#### 原理

布隆过滤器（Bloom Filter）的核心实现是一个超大的位数组和几个哈希函数。假设位数组的长度为m，哈希函数的个数为k

![配图](.\img\Snipaste_2021-06-03_23-41-54.png)

具体的操作流程：假设集合里面有3个元素{x, y, z}，哈希函数的个数为3。首先将位数组进行初始化，将里面每个位都设置位0。对于集合里面的每一个元素，将元素依次通过3个哈希函数进行映射，每次映射都会产生一个哈希值，这个值对应位数组上面的一个点，然后将位数组对应的位置标记为1。查询W元素是否存在集合中的时候，同样的方法将W通过哈希映射到位数组上的3个点。如果3个点的其中有一个点不为1，则可以判断该元素一定不存在集合中。反之，如果3个点都为1，则该元素可能存在集合中。注意：此处不能判断该元素是否一定存在集合中，可能存在一定的误判率。可以从图中可以看到：假设某个元素通过映射对应下标为4，5，6这3个点。虽然这3个点都为1，但是很明显这3个点是不同元素经过哈希得到的位置，因此这种情况说明元素虽然不在集合中，也可能对应的都是1，这是误判率存在的原因。

#### 布隆过滤器添加元素

- 将要添加的元素给k个哈希函数
- 得到对应于位数组上的k个位置
- 将这k个位置设为1

#### 布隆过滤器查询元素

- 将要查询的元素给k个哈希函数
- 得到对应于位数组上的k个位置
- 如果k个位置有一个为0，则肯定不在集合中
- 如果k个位置全部为1，则可能在集合中

#### 学习资料和示例代码
- 布隆过滤器的原理和实现 https://www.cnblogs.com/cpselvis/p/6265825.html
- 使用布隆过滤器解决缓存击穿、垃圾邮件识别、集合判重 https://blog.csdn.net/tianyaleixiaowu/article/details/74721877
- 布隆过滤器 Python 代码示例 https://shimo.im/docs/UITYMj1eK88JCJTH
- 布隆过滤器 Python 实现示例 https://www.geeksforgeeks.org/bloom-filters-introduction-and-python-implementation/
- 高性能布隆过滤器 Python 实现示例 https://github.com/jhgg/pybloof
- 布隆过滤器 Java 实现示例 1 https://github.com/lovasoa/bloomfilter/blob/master/src/main/java/BloomFilter.java
- 布隆过滤器 Java 实现示例 2 https://github.com/Baqend/Orestes-Bloomfilter

### LRU Cache - Least Rencent Used Cache

LRU 是一种十分常见的页面置换算法。将 LRU 翻译成大白话就是：当不得不淘汰某些数据时（通常是容量已满），选择最久未被使用的数据进行淘汰。

还有一些其他的页面替换算法，见 https://en.wikipedia.org/wiki/Cache_replacement_policies

#### 基本实现

LRU 缓存机制可以通过哈希表辅以双向链表实现，我们用一个哈希表和一个双向链表维护所有在缓存中的键值对。

- 双向链表按照被使用的顺序存储了这些键值对，靠近头部的键值对是最近使用的，而靠近尾部的键值对是最久未使用的。

- 哈希表即为普通的哈希映射（HashMap），通过缓存数据的键映射到其在双向链表中的位置。

这样以来，我们首先使用哈希表进行定位，找出缓存项在双向链表中的位置，随后将其移动到双向链表的头部，即可在 O(1) 的时间内完成 get 或者 put 操作。具体的方法如下：

- 对于 get 操作，首先判断 key 是否存在：
  - 如果 key 不存在，则返回 -1−1；
  - 如果 key 存在，则 key 对应的节点是最近被使用的节点。通过哈希表定位到该节点在双向链表中的位置，并将其移动到双向链表的头部，最后返回该节点的值。

- 对于 put 操作，首先判断 key 是否存在：
  - 如果 key 不存在，使用 key 和 value 创建一个新的节点，在双向链表的头部添加该节点，并将 key 和该节点添加进哈希表中。然后判断双向链表的节点数是否超出容量，如果超出容量，则删除双向链表的尾部节点，并删除哈希表中对应的项；
  - 如果 key 存在，则与 get 操作类似，先通过哈希表定位，再将对应的节点的值更新为 value，并将该节点移到双向链表的头部。

上述各项操作中，访问哈希表的时间复杂度为 O(1)，在双向链表的头部添加节点、在双向链表的尾部删除节点的复杂度也为 O(1)O。而将一个节点移到双向链表的头部，可以分成「删除该节点」和「在双向链表的头部添加节点」两步操作，都可以在 O(1)时间内完成。

小贴士

在双向链表的实现中，使用一个伪头部（dummy head）和伪尾部（dummy tail）标记界限，这样在添加节点和删除节点的时候就不需要检查相邻的节点是否存在。

#### 具体实现

见Leet Code题目 

- LRU 缓存机制 https://leetcode-cn.com/problems/lru-cache/


## 第九周 - 第十八课 - 排序算法

### Implementation For The Sort Algorithms In Golang

https://www.tutorialdocs.com/article/golang-sort-algorithms.html

### 各种排序算法动画展示

https://www.cnblogs.com/onepixel/p/7674659.html

#### O (N2) -- The comparison-based sort algorithms
- Bubble sort (BUB)
- Selection sort (SEL)
- Insert sort (INS)
  
The comparison-based sort algorithms compare the elements of the array and then decide whether to swap them. These three sorting algorithms are the easiest ones to implement, but are not the most efficient, because their time complexity is O (N2).

#### O (NlogN) -- The comparison-based sort algorithms
- Heap sort
- Shell Sort (SHE)
- Merge sort (MER)
- Quick sort (QUI)
- Random Quick Sort (R-Q)
- These sorting algorithms are usually implemented recursively.

#### O (N) -- The sort algorithms not based on comparison
- Counting sort (COU)
- Bucket sort (BUC)
- Radix Sort (RAD)

#### Complexity and stability

![配图](.\img\Snipaste_2021-06-06_00-10-56.png)

#### Bublle Sort 冒泡排序
```golang
func BubbleSort(data []int) {
	for i := 0; i < len(data); i++ {
		for j := 1; j < len(data)-i; j++ {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}
}
```