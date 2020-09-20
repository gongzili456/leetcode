package leetcode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) <= 0 || len(grid[0]) <= 0 {
		return 0
	}

	nx, ny := len(grid), len(grid[0])

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= nx || y < 0 || y >= ny || grid[x][y] != '1' {
			return
		}

		grid[x][y] = '0'
		dfs(x+1, y)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x, y-1)
	}

	ans := 0
	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}

// @lc code=end
