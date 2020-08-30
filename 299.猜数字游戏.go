package leetcode

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=299 lang=golang
 *
 * [299] 猜数字游戏
 */

// @lc code=start
func getHint(secret string, guess string) string {
	A, B := 0, 0
	S, G := make([]int, 10), make([]int, 10)

	for i := range secret {
		s, g := secret[i], guess[i]

		// 位置相同，则计数 A
		if s == g {
			A++
		} else {
			// 位置不同，分别在数组位置中 +1 计数
			S[s-'0']++
			G[g-'0']++
		}
	}

	// 遍历每一个计数位置，没有出现过则为 0，出现次数为最小的一方。
	for i := 0; i < 10; i++ {
		if S[i] >= G[i] {
			B += G[i]
		} else {
			B += S[i]
		}
	}

	return fmt.Sprintf("%dA%dB", A, B)
}

// @lc code=end
