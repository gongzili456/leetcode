package leetcode

/*
 * @lc app=leetcode.cn id=1528 lang=golang
 *
 * [1528] 重新排列字符串
 */

// @lc code=start
func restoreString(s string, indices []int) string {
	ans := make([]byte, len(s))

	for i := 0; i < len(s); i++ {
		ans[indices[i]] = s[i]
	}

	return string(ans)
}

// @lc code=end
