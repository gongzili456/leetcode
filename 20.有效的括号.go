package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s) <= 0 {
		return true
	}

	m := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []rune{}

	for _, r := range s {
		switch _, ok := m[r]; ok { // right side
		case true:
			if len(stack) == 0 || stack[len(stack)-1] != m[r] {
				return false
			}
			stack = stack[:len(stack)-1]
		case false:
			// left
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}

// @lc code=end
