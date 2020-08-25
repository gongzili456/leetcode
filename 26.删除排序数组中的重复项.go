package leetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	i := 0
	for j := 1; j < len(nums); j++ {
		// 只需要将不同的数字前移即可
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// @lc code=end
