package leetcode

/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[[2]int]bool, len(obstacles))
	for _, o := range obstacles {
		m[[2]int{o[0], o[1]}] = true
	}

	dirs := [4][2]int{
		{0, 1},  // up
		{1, 0},  // right
		{0, -1}, // down
		{-1, 0}, // left
	}

	var curDir, x, y, ans int

	for _, com := range commands {
		switch com {
		case -2:
			curDir = (curDir + 3) % 4
		case -1:
			curDir = (curDir + 1) % 4
		default:
			for i := 0; i < com; i++ {
				x += dirs[curDir][0]
				y += dirs[curDir][1]
				if _, ok := m[[2]int{x, y}]; ok {
					x -= dirs[curDir][0]
					y -= dirs[curDir][1]
				}
				ans = max(ans, x*x+y*y)
			}
		}
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end
