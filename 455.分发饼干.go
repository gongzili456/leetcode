package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	ans := 0
	j := 0

	for i := 0; i < len(g); i++ {
		for ; j < len(s) && s[j] < g[i]; j++ {
		}

		if j < len(s) {
			ans++
			j++
		}
	}

	return ans
}

// @lc code=end
