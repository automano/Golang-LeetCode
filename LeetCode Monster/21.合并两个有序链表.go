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
	// 2.迭代解法
	dummyHead := &ListNode{}
	tempNode := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tempNode.Next = l1
			l1 = l1.Next
		} else {
			tempNode.Next = l2
			l2 = l2.Next
		}
		tempNode = tempNode.Next
	}

	if l1 == nil {
		tempNode.Next = l2
	}
	if l2 == nil {
		tempNode.Next = l1
	}

	return dummyHead.Next
}

// @lc code=end

