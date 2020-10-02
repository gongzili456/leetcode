package leetcode

/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	myRob := func(ns []int) int {
		var pre, cur int
		for _, n := range ns {
			pre, cur = cur, max(pre+n, cur)
		}
		return cur
	}

	return max(myRob(nums[:len(nums)-1]), myRob(nums[1:]))
}

// @lc code=end
