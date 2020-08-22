package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	type band struct {
		i int    // integer
		n string // numeral
	}

	var bands = []band{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	res := 0

	// 与整数转罗马数字相似的逻辑
	for _, b := range bands {
		for strings.HasPrefix(s, b.n) {
			res += b.i
			s = s[len(b.n):]
		}
	}

	return res
}

// @lc code=end
