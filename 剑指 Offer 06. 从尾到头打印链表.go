package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	temp := []int{}

	current := head

	for current != nil {
		temp = append(temp, current.Val)
		current = current.Next
	}

	ans := []int{}
	for len(temp) > 0 {
		n := temp[len(temp)-1]
		ans = append(ans, n)
		temp = temp[:len(temp)-1]
	}

	return ans
}
