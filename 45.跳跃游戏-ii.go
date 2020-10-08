package leetcode

/*
 * @lc app=leetcode.cn id=45 lang=golang
 *
 * [45] 跳跃游戏 II
 */

// @lc code=start
func jump(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])

		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}

// @lc code=end
