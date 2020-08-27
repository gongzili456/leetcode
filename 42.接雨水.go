package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	s := []int{0}
	current, ans := 0, 0

	for current < len(height) {
		for len(s) > 0 && height[current] > height[s[len(s)-1]] {
			top := s[len(s)-1]
			s = s[:len(s)-1]

			if len(s) <= 0 {
				break
			}

			distance := current - top - 1
			h := s[len(s)-1]
			if height[current] < h {
				h = height[current]
			}
			h -= height[top]

			ans += distance * h
		}

		current++
		s = append(s, current)
	}

	return ans
}

// @lc code=end
