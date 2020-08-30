package leetcode

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	str := []byte{}

	left := 0
	for i := range s {
		if s[i] != ' ' {
			continue
		}

		for j := i - 1; j >= left; j-- {
			str = append(str, s[j])
		}
		str = append(str, s[i])
		left = i + 1
	}

	if left < len(s) {
		for j := len(s) - 1; j >= left; j-- {
			str = append(str, s[j])
		}
	}

	return string(str)
}

// @lc code=end
