/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 2. 迭代
	// 哨兵节点 dummyHead
	dummyHead := &ListNode{}
	temp := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			temp.Next = l1 // 构建合并链表
			l1 = l1.Next 	// 把已被连接的节点移除掉
		} else {
			temp.Next = l2
			l2 = l2.Next
		}
		// 移动temp
		temp = temp.Next		
	}
	if l1 == nil {
		temp.Next = l2
	}
	if l2 == nil {
		temp.Next = l1
	}
	return dummyHead.Next
}
// @lc code=end

