package leetcode

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
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
func sumOfLeftLeaves(root *TreeNode) int {
	var dfs func(n *TreeNode) int
	dfs = func(node *TreeNode) (ans int) {
		if node.Left != nil {
			if isLeafNode(node.Left) {
				ans += node.Left.Val
			} else {
				ans += dfs(node.Left)
			}
		}
		if node.Right != nil && !isLeafNode(node.Right) {
			ans += dfs(node.Right)
		}
		return
	}

	if root == nil {
		return 0
	}
	return dfs(root)
}

func isLeafNode(n *TreeNode) bool {
	return n.Left == nil && n.Right == nil
}

// @lc code=end
