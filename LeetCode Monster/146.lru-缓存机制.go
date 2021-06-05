/*
 * @lc app=leetcode.cn id=146 lang=golang
 *
 * [146] LRU 缓存机制
 */

// @lc code=start
type LRUCache struct {
	size       int                  // 当前元素个数
	capacity   int                  // 缓存容量
	cache      map[int]*DLinkedNode // 缓存单元
	head, tail *DLinkedNode         // 哨兵节点
}

type DLinkedNode struct {
	key, val   int          // 缓存单元key：value
	prev, next *DLinkedNode // 双向链表指针
}

func initDLinkedNode(key, val int) *DLinkedNode {
	// 初始化双向链表节点
	return &DLinkedNode{
		key: key,
		val: val,
	}
}

func Constructor(capacity int) LRUCache {
	// 初始化缓存
	l := LRUCache{
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	// 如果key不存在，返回-1
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	// 找到对应key的node
	node := this.cache[key]
	// node被访问，移到链表头
	this.moveToHead(node)
	return node.val
}

func (this *LRUCache) Put(key int, value int) {

	if _, ok := this.cache[key]; !ok {
		// key 不存在，将新node放入
		node := initDLinkedNode(key, value)
		// 新node放入缓存
		this.cache[key] = node
		// 新node加入链表头
		this.addToHead(node)
		this.size++
		// 判断是否超出容量
		if this.size > this.capacity {
			// 从链表尾删除元素
			removed := this.removeTail()
			// 释放该节点占用的哈希表位置
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		// key已存在，更新其值
		node := this.cache[key]
		node.val = value
		// node被访问，移到链表头
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	// 更新node的前后指针
	node.prev = this.head
	node.next = this.head.next
	// 更新指向node的指针
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	// 当前位置删除该节点
	this.removeNode(node)
	// 添加到链表头
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	// 通过哨兵tail获取最后一个元素
	node := this.tail.prev
	// 删除该节点
	this.removeNode(node)
	return node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

