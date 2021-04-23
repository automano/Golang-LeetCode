/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	result := []int{}
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		queueSize := len(queue)
		maxTemp := math.MinInt64
		for i := 0; i < queueSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Val > maxTemp {
				maxTemp = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, maxTemp)
	}
	return result
}

// @lc code=end

