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
//  递归法
// func inorderTraversal(root *TreeNode) []int {
// 	if root == nil {
// 		return []int{}
// 	}

// 	// 左跟右
// 	ans := []int{}

// 	var helper func(node *TreeNode)
// 	helper = func(node *TreeNode) {
// 		if node.Left != nil {
// 			helper(node.Left)
// 		}

// 		ans = append(ans, node.Val)

// 		if node.Right != nil {
// 			helper(node.Right)
// 		}
// 	}

// 	helper(root)

// 	return ans
// }

// 辅助栈
func inorderTraversal(root *TreeNode) []int {
	stack := []*TreeNode{}
	ans := []int{}

	current := root
	for current != nil || len(stack) > 0 {
		// 将所有左子节点压栈
		for current != nil {
			stack = append(stack, current)
			current = current.Left
		}

		current = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, current.Val)

		current = current.Right
	}
	return ans
}

// @lc code=end
