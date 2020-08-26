package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	p1, p2 := m-1, n-1

	for p1 >= 0 || p2 >= 0 {
		switch {
		case p1 < 0:
			nums1[p] = nums2[p2]
			p--
			p2--
		case p2 < 0:
			nums1[p] = nums1[p1]
			p--
			p1--
		default:
			if nums1[p1] <= nums2[p2] {
				nums1[p] = nums2[p2]
				p2--
			} else {
				nums1[p] = nums1[p1]
				p1--
			}
			p--
		}
	}
}

// @lc code=end
