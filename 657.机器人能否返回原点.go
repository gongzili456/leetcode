package leetcode

/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	x, y := 0, 0

	for _, r := range moves {
		switch r {
		case 'R':
			x++
		case 'L':
			x--
		case 'D':
			y--
		case 'U':
			y++
		}
	}
	return x == 0 && y == 0
}

// @lc code=end
