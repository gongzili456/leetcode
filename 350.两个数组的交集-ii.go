package leetcode

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	dic := map[int]int{}

	for _, v := range nums1 {
		dic[v]++
	}

	ans := []int{}

	for i := 0; i < len(nums2); i++ {
		if v, ok := dic[nums2[i]]; ok && v > 0 {
			dic[nums2[i]]--
			ans = append(ans, nums2[i])
		}
	}

	return ans
}

// @lc code=end
