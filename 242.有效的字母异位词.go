package leetcode

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	// 如果两个字符串包含的字符和数量都相同，则是异位词
	if len(s) != len(t) {
		return false
	}

	m := map[rune]int{}

	for _, r := range s {
		m[r]++
	}

	for _, r := range t {
		v, ok := m[r]
		if !ok || v == 0 {
			return false
		}
		m[r]--
	}

	return true
}

// @lc code=end
