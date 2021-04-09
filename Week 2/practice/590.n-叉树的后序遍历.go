/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N 叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	// 1. 递归写法
	result := []int{}
	var postorderTerversal func(node *Node)
	postorderTerversal = func(node *Node) {
		if node == nil {
			return
		}
		for i := 0; i < len(node.Children); i++ {
			postorderTerversal(node.Children[i])
		}
		result = append(result, node.Val)
	}
	postorderTerversal(root)
	return result
}

// @lc code=end

