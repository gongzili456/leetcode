package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	// 排序可以保证相邻的重复数字只使用一次
	sort.Ints(nums)
	n := len(nums)
	ans := [][]int{}
	perm := []int{}
	vis := make([]bool, n)

	var backtrack func(int)
	backtrack = func(idx int) {
		if idx == n {
			ans = append(ans, append([]int(nil), perm...))
			return
		}

		for i, v := range nums {
			// 如果已经访问，或与前面数字相同
			if vis[i] || i > 0 && !vis[i-1] && v == nums[i-1] {
				continue
			}

			perm = append(perm, v)
			vis[i] = true
			backtrack(idx + 1)
			vis[i] = false
			perm = perm[:len(perm)-1]
		}
	}

	backtrack(0)

	return ans
}

// @lc code=end
