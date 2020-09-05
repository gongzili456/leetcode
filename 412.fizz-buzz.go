package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	ans := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			ans = append(ans, "FizzBuzz")
			continue
		}

		if i%3 == 0 {
			ans = append(ans, "Fizz")
			continue
		}

		if i%5 == 0 {
			ans = append(ans, "Buzz")
			continue
		}

		ans = append(ans, fmt.Sprintf("%d", i))
	}

	return ans
}

// @lc code=end
