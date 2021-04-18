/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	// 3. BFS 层序遍历
	if root == nil {
		return 0
	}
	maxResult := 0
	queue := []*TreeNode{root} // 用队列保存中间节点

	for len(queue) != 0 {
		queueSize := len(queue)

		for queueSize > 0 {
			// 用数组模拟队列的先进先出
			node := queue[0]
			queue = queue[1:len(queue)]
			// 子节点压入队列中
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queueSize--
		}
		maxResult++
	}
	return maxResult
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end

