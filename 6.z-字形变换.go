package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	ret := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		ret[i] = make([]byte, 0)
	}

	i, dir := 0, -1

	for j := 0; j < len(s); j++ {
		ret[i] = append(ret[i], s[j])

		// 首行或者尾行进行转向
		if i == 0 || i == numRows-1 {
			dir = -dir
		}
		i += dir
	}

	res := []string{}

	for _, re := range ret {
		res = append(res, string(re))
	}

	return strings.Join(res, "")
}

// @lc code=end
