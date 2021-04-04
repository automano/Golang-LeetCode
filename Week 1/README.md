# 学习总结
## 如何高效学习
- 三分看视频理解，七分靠刷题练习
- 坚持反复做题（五毒神掌 - 五步刷题法）
- 记忆知识点，使用脑图
- 抛弃旧的习惯，不要死磕，敢于放手，敢于死记硬背（参考别人高赞答案)
- 最大误区：LeetCode题目只做一遍！

## 刷题方法
五毒神掌 - 五遍刷题法 (切记题目只做一遍！！！)
1. 第一遍：
   - 5～15分钟思考，不会马上看高赞题解，选择比较好的题解，先记住，默写下来

2. 第二遍：
   - 不看别人的代码，马上自己写下来，进行debug（自己开一个浏览器，空白的文件写代码）
   - 写完之后->LeetCode提交，直到通过的状态
   - 一个题目的话有多种的解法，多去体会。--->优化

3. 第三遍：
   - 过了24小时之后，回过头来做重复的题目
   - 不同解法的熟练程度--->专项练习

4. 第四遍：
   - 过了一周之后：反复回来进行练习相同的题目

5. 第五遍：
   - 面试前一周恢复性训练

## 时间复杂度
### 定义
   - O(1): Constant Complexity 常数复杂度
   - O(log n): Logarithmic Complexity 对数复杂度
   - O(n): Linear Complexity 线性时间复杂度
   - O(n^2): N square Complexity 平方
   - O(n^3): N cubic Complexity 立方
   - O(2^n): Exponential Growth 指数
   - O(n!): Factorial 阶乘
### 特殊
   - 二叉树遍历 - 前序、中序、后序：O(N)
   - 图的遍历：O(N)
   - 搜索算法：DFS、BFS - O(N)
   - 二分查找：O(logN)

## 第3课 
### 数组

   https://golang.google.cn/ref/spec#Array_types

   https://golang.google.cn/ref/spec#Slice_types

   - Array 在 golang中 a := [2]int 长度不可变，数组在内存中具有连续的地址
   - 更多的是使用切片，切片是一个结构体，（长度，容量，指向底层数组的指针），类似一个“可变”数组 a := make([]int,2)
### 链表

golang中官方提供的package中有一个对双向链表的实现 

https://golang.google.cn/pkg/container/list/

prepend O(1)

append O(1)

lookup O(n)

insert O(1)

delete O(1)

LeetCode设计链表 https://leetcode-cn.com/problems/design-linked-list/

### 跳表
利用索引，一种升级维度的思想，用空间换时间

- 只能用于元素有序的情况。
- 对标的是平衡树（AVL Tree）和二分查找，是一种 插入/删除/搜索 都是 O(log n) 的数据结构
- 应用于Redis

跳表简介 https://www.jianshu.com/p/fef9817cc943

LeetCode设计跳表 https://leetcode-cn.com/problems/design-skiplist/

## 第4课
### 栈
- Stack：先入后出；添加、删除皆为 O(1)
- 栈是一个 LIFO 数据结构。通常，插入操作在栈中被称作入栈 push 。与队列类似，总是在堆栈的末尾添加一个新元素。但是，删除操作，退栈 pop ，将始终删除队列中相对于它的最后一个元素

LeetCode用栈实现队列 https://leetcode-cn.com/problems/implement-queue-using-stacks/
  
### 队列
- Queue：先入先出；添加、删除皆为 O(1)

LeetCode 设计循环队列 https://leetcode-cn.com/problems/design-circular-queue/

LeetCode 用队列实现栈 https://leetcode-cn.com/problems/implement-stack-using-queues/
### 双端队列

- 两端可以进出的 Queue Deque - double ended queue
- 插入和删除都是 O(1) 操作
  
### 优先队列

golang 使用官方包container/heap实现优先队列
https://golang.google.cn/pkg/container/heap/
