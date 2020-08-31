package leetcode

/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	vis := make([]bool, len(rooms))
	total := 0

	var dfs func(x int)
	dfs = func(x int) {
		vis[x] = true
		total++

		for _, r := range rooms[x] {
			if !vis[r] {
				dfs(r)
			}
		}
	}

	dfs(0)
	return total == len(rooms)
}

// @lc code=end
