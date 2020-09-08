package leetcode

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
// 解法：
// 递归
// 利用队列
// 广度优先搜索
func levelOrder(root *Node) [][]int {
	ans := [][]int{}

	var helper func(n *Node, level int)
	helper = func(n *Node, level int) {
		if n == nil {
			return
		}

		if len(ans) == level {
			ans = append(ans, []int{})
		}

		ans[level] = append(ans[level], n.Val)

		for _, c := range n.Children {
			helper(c, level+1)
		}
	}

	helper(root, 0)
	return ans
}

// @lc code=end
