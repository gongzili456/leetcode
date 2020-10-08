package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	n := len(nums)
	rightmost := 0

	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}

	for i := 0; i < n; i++ {
		if i <= rightmost {
			// 可以到达的最右面的下标位置
			rightmost = max(rightmost, i+nums[i])

			if rightmost >= n-1 {
				return true
			}
		}
	}

	return false
}

// @lc code=end
