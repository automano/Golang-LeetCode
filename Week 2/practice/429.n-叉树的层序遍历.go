/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		temp := make([]int, 0)
		queueSize := len(queue)
		for i := 0; i < queueSize; i++ {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			for j := 0; j < len(node.Children); j++ {
				queue = append(queue, node.Children[j])
			}

		}
		result = append(result, temp)
	}
	return result
}

// @lc code=end

