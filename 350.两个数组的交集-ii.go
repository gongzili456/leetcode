package leetcode

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	dic := map[int]int{}
	ans := []int{}

	for _, n := range nums1 {
		dic[n]++
	}

	for _, n := range nums2 {
		if v, ok := dic[n]; ok && v > 0 {
			dic[n]--
			ans = append(ans, n)
		}
	}

	return ans
}

// @lc code=end
