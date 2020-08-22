package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 滑动窗口
	m, left, max := map[rune]int{}, 0, 0

	for i, r := range s {
		if _, ok := m[r]; ok && m[r] >= left {
			if i-left > max {
				max = i - left
			}
			left = m[r] + 1
		}

		m[r] = i
	}

	if len(s)-left > max {
		max = len(s) - left
	}

	return max
}

// @lc code=end
