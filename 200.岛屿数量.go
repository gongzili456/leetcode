package leetcode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr, nc := len(grid), len(grid[0])
	islands := 0

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= nr || c < 0 || c >= nc || grid[r][c] == '0' {
			return
		}

		grid[r][c] = '0'
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				islands++
				dfs(r, c)
			}
		}
	}

	return islands
}

// @lc code=end
