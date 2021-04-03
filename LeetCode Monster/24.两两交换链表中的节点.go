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
	// 递归解法
	// 1. 返回值：完成交换的子链表 （除去当前正在考虑的两个节点）
	// 2. 调用单元： 当前的组合 head和next. head -> 子链表， next -> head 完成了当前单元的交换
	// 3. 终止条件： head -> null 没有节点了 head.Next -> null 奇数个节点的情况
	
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	// 开始当前单元的交换
	head.Next = swapPairs(next.Next)
	next.Next = head
	// 完成当前单元的交换
	return next 
}
// @lc code=end

