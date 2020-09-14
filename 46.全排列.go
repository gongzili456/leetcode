package leetcode

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ans := [][]int{}
	visied := make([]bool, len(nums))
	temp :=  []int{}

	// 深度优先搜索 + 回溯
	var dfs func(i int)
	dfs = func(i int){
		if i == len(nums) {
			ans = append(ans, append([]int(nil), temp...))
			return
		}

		for j, n := range nums {
			if !visied[j] {
					// 标记并添加至结果集
					visied[j] = true
					temp = append(temp, n)

					dfs(i+1)

					// 移出标记，并移出结果集
					visied[j] = false
					temp = temp[:len(temp)-1]
				}	
		}
	}

	dfs(0)
	return ans
}
// @lc code=end

