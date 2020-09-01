package leetcode

/*
 * @lc app=leetcode.cn id=486 lang=golang
 *
 * [486] 预测赢家
 */

// PredictTheWinner is
// @lc code=start
func PredictTheWinner(nums []int) bool {
	var rec func(start, end int) int
	rec = func(start, end int) int {
		if start == end {
			return nums[start]
		}

		// 当前得分与之后总得分的差，取大值的情况返回
		left := nums[start] - rec(start+1, end)
		right := nums[end] - rec(start, end-1)

		if left > right {
			return left
		}

		return right
	}

	return rec(0, len(nums)-1) >= 0
}

// @lc code=end
