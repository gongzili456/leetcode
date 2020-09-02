package leetcode

/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(S string) string {
	ans, stack := []rune{}, []rune{}

	for _, s := range S {
		// 当前为（ 时，且栈为空，则是左边界
		if s == '(' {
			if len(stack) > 0 {
				ans = append(ans, s)
			}
			stack = append(stack, s)
		}

		// 当前为 ）时，出栈后，栈为空，则是右边界
		if s == ')' {
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				ans = append(ans, s)
			}
		}
	}

	return string(ans)
}

// @lc code=end
