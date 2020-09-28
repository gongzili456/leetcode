package leetcode

/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// Node is
// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	q := []*Node{root}

	for len(q) > 0 {
		// 逐层遍历
		tmp := q
		q = nil

		for i, node := range tmp {
			// 处理非末尾节点
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return root
}

// @lc code=end
