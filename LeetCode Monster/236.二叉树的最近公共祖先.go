/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归解法
	//终止条件
	if root == nil || root == p || root == q {
		return root
	}
	// 处理当前层
	// drill down
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 返回值
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

// @lc code=end

