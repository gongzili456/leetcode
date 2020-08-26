package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := map[int]int{}

	for i, n := range nums {
		t := target - n
		if p, ok := m[t]; ok {
			return []int{p, i}
		}
		m[n] = i
	}

	return nil
}

// @lc code=end
