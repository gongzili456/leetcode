package leetcode

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	dp := make([]int, amount+1)

	// TODO: 还没懂
	for target := range dp {
		for _, coin := range coins {
			if target == coin {
				dp[target] = 1
				continue
			}

			left := target - coin
			if left >= 0 && dp[left] != 0 {
				if dp[target] != 0 {
					dp[target] = min(dp[target], dp[left]+1)
				} else {
					dp[target] = dp[left] + 1
				}
			}
		}
	}

	if dp[amount] == 0 {
		return -1
	}

	return dp[amount]
}

// @lc code=end
