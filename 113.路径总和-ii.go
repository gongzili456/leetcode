package leetcode

/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	ans := [][]int{}
	path := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}

		left -= node.Val
		path = append(path, node.Val)

		defer func() {
			path = path[:len(path)-1]
		}()

		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}

		dfs(node.Left, left)
		dfs(node.Right, left)
	}

	dfs(root, sum)
	return ans
}

// @lc code=end
