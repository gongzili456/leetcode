package leetcode

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// Node is
type Node struct {
	Val      int
	Children []*Node
}

// @lc code=start
/**
 * Definition for a Node.
 */
func preorder(root *Node) []int {
	ans := []int{}

	var front func(n *Node)
	front = func(n *Node) {
		if n == nil {
			return
		}

		ans = append(ans, n.Val)

		for _, c := range n.Children {
			front(c)
		}
	}

	front(root)
	return ans
}

// @lc code=end

//
