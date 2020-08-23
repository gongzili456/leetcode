package leetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	left, right := '(', ')'
	ans := []string{}

	var backtrack func(chars []rune, l, r int)
	backtrack = func(chars []rune, l, r int) {
		if len(chars) == n*2 {
			ans = append(ans, string(chars))
			return
		}

		// 左边可以随便放，只要不超量
		if l < n {
			backtrack(append(chars, left), l+1, r)
		}

		// 右边必须小于左边的数量
		if r < l {
			backtrack(append(chars, right), l, r+1)
		}
	}

	backtrack(nil, 0, 0)
	return ans
}

// @lc code=end
