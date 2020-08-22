package leetcode

import "strings"

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := []string{}

	// Loop through each symbol, stopping if num becomes 0.
	for i := 0; i < len(values) && num > 0; i++ {
		// Repeat while the current symbol still fits into num.
		for values[i] <= num {
			num -= values[i]
			res = append(res, symbols[i])
		}
	}

	return strings.Join(res, "")
}

// @lc code=end
