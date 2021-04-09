/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	// 1. 递归写法
	result := []int{}
	var preorderTerversal func(node *Node)
	preorderTerversal = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for i := 0; i < len(node.Children); i++ {
			preorderTerversal(node.Children[i])
		}
	}
	preorderTerversal(root)
	return result
}

// @lc code=end

