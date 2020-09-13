package leetcode

/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */

// @lc code=start
func combinationSum3(k int, n int) [][]int {
	ans := [][]int{}
	temp := []int{}

	var dfs func(curr, rest int)
	dfs = func(curr, rest int) {
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}

		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-curr < k || rest < 0 {
			return
		}

		dfs(curr+1, rest)

		temp = append(temp, curr)
		dfs(curr+1, rest-curr)
		temp = temp[:len(temp)-1]
	}

	dfs(1, n)
	return ans
}

// @lc code=end
