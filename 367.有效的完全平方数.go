package leetcode

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	if num <= 2 {
		return true
	}

	left, right := 0, num/2

	for left < right {
		x := (left + right) / 2
		if x*x == num {
			return true
		}

		if x*x > num {
			right = x - 1
		} else {
			left = x + 1
		}
	}

	return false
}

// @lc code=end
