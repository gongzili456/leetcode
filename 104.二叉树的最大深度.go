package leetcode

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
	// 抵达叶子节点，没有子节点，不计入深度
	if root == nil {
		return 0
	}

	// 计算左侧深度
	left := maxDepth(root.Left)

	// 计算右侧深度
	right := maxDepth(root.Right)

	// 选择更深的一边
	if left < right {
		left = right
	}

	// 返回更深的一边，加上自身的深度 1
	return left + 1
}

// @lc code=end
