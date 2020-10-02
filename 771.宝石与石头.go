package leetcode

/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	// 也可以用数组
	m := map[byte]int{}

	for i := 0; i < len(J); i++ {
		m[J[i]] = 0
	}

	count := 0

	for i := 0; i < len(S); i++ {
		if _, ok := m[S[i]]; ok {
			count++
		}
	}

	return count
}

// @lc code=end
