package leetcode

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	a := 1

	for i := len(digits) - 1; i >= 0; i-- {
		a += digits[i]
		digits[i] = a % 10
		a /= 10
	}

	if a > 0 {
		digits = append([]int{a}, digits...)
	}

	return digits
}

// @lc code=end
