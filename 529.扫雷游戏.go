package leetcode

/*
 * @lc app=leetcode.cn id=529 lang=golang
 *
 * [529] 扫雷游戏
 */

// @lc code=start
func updateBoard(board [][]byte, click []int) [][]byte {
	dirX := []int{0, 1, 0, -1, 1, 1, -1, -1}
	dirY := []int{1, 0, -1, 0, 1, -1, 1, -1}

	var dfs func(x, y int)
	dfs = func(x, y int) {
		cnt := 0

		// 检查当前位置周边8个位置是否有地雷
		for i := 0; i < 8; i++ {
			tx, ty := x+dirX[i], y+dirY[i]
			if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) {
				continue
			}

			if board[tx][ty] == 'M' {
				cnt++
			}
		}

		if cnt > 0 {
			board[x][y] = byte(cnt + '0')
		} else {
			board[x][y] = 'B'

			// 递进下一层搜索
			for i := 0; i < 8; i++ {
				tx, ty := x+dirX[i], y+dirY[i]

				if tx < 0 || tx >= len(board) || ty < 0 || ty >= len(board[0]) || board[tx][ty] != 'E' {
					continue
				}
				dfs(tx, ty)
			}
		}
	}

	x, y := click[0], click[1]
	if board[x][y] == 'M' {
		board[x][y] = 'X'
	} else {
		dfs(x, y)
	}

	return board
}

// @lc code=end
