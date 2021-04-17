/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 前序遍历 根->左->右
	// 中序遍历 左->根->右
	if len(preorder) == 0 {
		return nil
	}
	// 初始化root节点 -> 前序遍历结果的第一个元素
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	// 在中序遍历结果中找到根节点的位置
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}

// @lc code=end

