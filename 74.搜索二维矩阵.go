package leetcode

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])

	left, right := 0, m*n-1
	var pivotIdx, pivotElement int

	for left <= right {
		pivotIdx = (left + right) / 2
		pivotElement = matrix[pivotIdx/n][pivotIdx%n]

		if target == pivotElement {
			return true
		}

		if target < pivotElement {
			right = pivotIdx - 1
		} else {
			left = pivotIdx + 1
		}
	}

	return false
}

// @lc code=end
