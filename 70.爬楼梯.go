package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	// 转换成斐波那契数列问题
	// f(n) = f(n-1) + f(n-2)
	if n == 1 {
		return 1
	}

	first, second := 1, 2

	for i := 3; i <= n; i++ {
		third := first + second
		first, second = second, third
	}

	return second
}

// @lc code=end
