package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	// four cases check
	// 1. discards all leading whitespaces
	// 2. sign of the number
	// 3. overflow
	// 4. invalid input
	sign, base, i, n := 1, 0, 0, len(str)

	for i < n && str[i] == ' ' {
		i++
	}

	if i < n {
		if str[i] == '-' {
			sign = -1
			i++
		} else if str[i] == '+' {
			i++
		}
	}

	for i < n && str[i] >= '0' && str[i] <= '9' {
		base = base*10 + int(str[i]) - '0'

		if sign*base > math.MaxInt32 {
			return math.MaxInt32
		} else if sign*base < math.MinInt32 {
			return math.MinInt32
		}

		i++
	}

	return base * sign
}

// @lc code=end
