package leetcode

/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num < 10 {
		return num
	}

	next := 0
	for num != 0 {
		// 从个位数逐个相加
		next += num % 10
		num /= 10
	}

	return addDigits(next)
}

// @lc code=end
