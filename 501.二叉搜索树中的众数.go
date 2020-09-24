package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=501 lang=golang
 *
 * [501] 二叉搜索树中的众数
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
func findMode(root *TreeNode) []int {
	var base, count, maxCount int
	answer := []int{}

	update := func(x int) {
		if x == base {
			count++
		} else {
			base = x
			count = 1
		}

		// 如果出现的次数相等
		if count == maxCount {
			answer = append(answer, base)
		} else if count > maxCount {
			// 新的众数
			maxCount = count
			answer = []int{base}
		}
	}

	cur := root
	for cur != nil {
		fmt.Println("cur: ", cur.Val)
		if cur.Left == nil {
			update(cur.Val)
			cur = cur.Right
			continue
		}

		// TODO: 需要继续阅读题解
		pre := cur.Left
		for pre.Right != nil && pre.Right != cur {
			pre = pre.Right
		}

		if pre.Right == nil {
			pre.Right = cur
			cur = cur.Left
		} else {
			pre.Right = nil
			update(cur.Val)
			cur = cur.Right
		}
	}

	return answer
}

// @lc code=end
