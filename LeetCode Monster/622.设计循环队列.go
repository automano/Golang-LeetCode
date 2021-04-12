/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	data  []int
	front int
	rear  int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		data:  make([]int, k+1),
		front: 0,
		rear:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[this.rear] = value
	this.rear = (this.rear + 1) % len(this.data)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.front = (this.front + 1) % len(this.data)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.rear-1+len(this.data))%len(this.data)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.front == this.rear
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.rear+1)%len(this.data) == this.front
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

