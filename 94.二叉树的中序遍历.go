package leetcode

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// TreeNode is
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	// 中序，左跟右
	stack := []*TreeNode{}
	ans := []int{}

	for root != nil ||  len(stack) > 0 {
		// 每次都要先检查当前节点所有的左子树
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		// 然后再处理跟节点
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, root.Val)

		// 然后再处理右子树
		root = root.Right
	}

	return ans
}

func inorderTraversal2(root *TreeNode) []int {
	// 中序，左跟右
	ans := []int{}

	var helper func(r *TreeNode)
	helper = func (r *TreeNode)  {
		if r == nil {
			return
		}

		helper(r.Left)
		ans = append(ans, r.Val)
		helper(r.Right)
	}

	helper(root)
	return ans
}

// @lc code=end
