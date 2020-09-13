package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) (ans [][]int) {
	// 预先排序，方便之后剪枝
	sort.Ints(candidates)

	// 统计不重复的数字和出现的次数，犹豫是已经排序了，所以 freq[i][0] 也是有序的
	freq := [][2]int{}
	for _, n := range candidates {
		if len(freq) == 0 || n != freq[len(freq)-1][0] {
			freq = append(freq, [2]int{n, 1})
		} else {
			freq[len(freq)-1][1]++
		}
	}

	temp := []int{}

	var dfs func(pos, sum int)
	dfs = func(pos, sum int) {
		if sum == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}

		// 如果 sum 小于后面的数字，则无法得到 target，不用计算了，因为之前已经排序了
		if pos == len(freq) || sum < freq[pos][0] {
			return
		}

		// 跳过当前数字
		dfs(pos+1, sum)

		// sum / freq[pos][0] 是最多能用的数量，freq[pos][1] 是最多可用的数量
		most := min(sum/freq[pos][0], freq[pos][1])

		for i := 1; i <= most; i++ {
			temp = append(temp, freq[pos][0])
			dfs(pos+1, sum-freq[pos][0]*i)
		}

		// 回溯
		temp = temp[:len(temp)-most]
	}

	dfs(0, target)
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
