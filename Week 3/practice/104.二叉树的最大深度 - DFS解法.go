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
	// 2. DFS 迭代
	if root == nil {
		return 0
	}
	maxResult := 0
	stack := []*TreeNode{root} // 用栈保存中间节点
	depth := []int{1}          // 层数
	for len(stack) != 0 {
		// 用数组模拟栈的先进后出
		node := stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]
		// 当前层深度
		temp := depth[len(depth)-1]
		depth = depth[0 : len(depth)-1]
		// 维护最大层数
		maxResult = max(maxResult, temp)
		// 子节点压入栈中
		// 层数+1 压入栈中
		if node.Left != nil {
			stack = append([]*TreeNode{node.Left}, stack...)
			depth = append([]int{temp + 1}, depth...)
		}
		if node.Right != nil {
			stack = append([]*TreeNode{node.Right}, stack...)
			depth = append([]int{temp + 1}, depth...)
		}
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

