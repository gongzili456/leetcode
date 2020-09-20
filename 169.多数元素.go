package leetcode

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	max := 0
	m := map[int]int{}
	target := len(nums) / 2
	for _, n := range nums {
		m[n]++
		v := m[n]
		if v > target && v > max {
			max = n
		}
	}

	return max
}

// @lc code=end
