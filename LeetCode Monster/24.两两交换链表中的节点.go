/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 迭代解法
	// 初始化 dummyHead -> head  (node1)
	// 交换时 temp -> node2 (node1.Next)
	//       node1 -> node3 (node2.Next)
	//       node2 -> node1
	// 移动temp == node1
	// 终止 temp -> nil 或者temp.Next -> nil
	// 返回 dummyHead.Next
	dummyHead := &ListNode{0,head} // 如何初始化一个结构体
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		// 定义交换单元
		node1 := temp.Next
		node2 := temp.Next.Next
		// 开始交换
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		// 移动temp
		temp = node1
	}
	return dummyHead.Next
}
// @lc code=end

