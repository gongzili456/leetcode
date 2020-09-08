package leetcode

/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	temp := []int{}
	ans := [][]int{}

	// 1, 2, 3, 4
	var dfs func(curr int)
	dfs = func(curr int) {
		if len(temp)+(n-curr+1) < k {
			return
		}

		if len(temp) == k {
			// 创建新数组，避免影响数据
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}

		// use
		temp = append(temp, curr)
		dfs(curr + 1)

		// unuse
		temp = temp[:len(temp)-1]
		dfs(curr + 1)
	}

	dfs(1)
	return ans
}

// @lc code=end
