package leetcode

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
	// 前序：跟左右
	// 中序：左跟右
	if len(preorder) == 0 {
		return nil
	}

	// 前序遍历的 0 索引位置即整棵树的根节点
	root := &TreeNode{preorder[0], nil, nil}
	i := 0

	// 找到中序的根节点索引位置
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
