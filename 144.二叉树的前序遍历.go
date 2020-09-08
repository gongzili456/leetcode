package leetcode

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) []int {
	ans := []int{}

	var helper func(n *TreeNode)
	helper = func(n *TreeNode) {
		if n == nil {
			return
		}

		ans = append(ans, n.Val)

		helper(n.Left)
		helper(n.Right)
	}

	helper(root)
	return ans
}

// @lc code=end
