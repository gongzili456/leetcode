package leetcode

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func swapPairs(head *ListNode) *ListNode {
// 	// 递归
// 	if head == nil || head.Next == nil {
// 		return head
// 	}

// 	first, second := head, head.Next

// 	// 交换就是，1 -> 3, 2 -> 1
// 	// 下一轮交换
// 	first.Next = swapPairs(second.Next)
// 	second.Next = first

// 	return second
// }

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	// 迭代
	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy

	for head != nil && head.Next != nil {
		first, second := head, head.Next

		pre.Next = second
		first.Next = second.Next
		second.Next = first

		pre = first
		head = first.Next
	}

	return dummy.Next
}

// @lc code=end
