package leetcode

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	temp := []int{}

	// sum: 剩余总和， i：数组当前下标
	var dfs func(sum, i int)
	dfs = func(sum, i int) {
		// 不能超出数组总长
		if i == len(candidates) {
			return
		}

		// 剩余总和为零，则 temp 中是元素和为 target
		if sum == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}

		// 不使用当前元素
		dfs(sum, i+1)

		// 如果还能还有剩余 sum 需要计算
		if sum-candidates[i] >= 0 {
			// 将当前元素加入队列
			temp = append(temp, candidates[i])

			// 因为可以重复使用元素，所以传入当前元素继续计算
			dfs(sum-candidates[i], i)

			// 最后，计算完成后将当前元素出队
			temp = temp[:len(temp)-1]
		}
	}

	dfs(target, 0)
	return ans
}

// @lc code=end
