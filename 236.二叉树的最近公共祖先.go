package leetcode

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 当前节点为空，或是目标节点是，停止深入递归
	if root == nil || root == p || root == q {
		return root
	}

	// 检查左子树
	left := lowestCommonAncestor(root.Left, p, q)

	// 检查右子树
	right := lowestCommonAncestor(root.Right, p, q)

	// 左侧为空，则目标可能在右边
	if left == nil {
		return right
	}

	// 右侧为空，则目标可能在左边
	if right == nil {
		return left
	}

	// 两侧都不为空，说明目标节点，分布在当前节点的左右子树上，自己就是最近祖先节点
	return root
}

// @lc code=end
