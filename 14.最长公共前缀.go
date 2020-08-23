package leetcode

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	// 二分查找
	if len(strs) <= 0 {
		return ""
	}

	isCommonPrefix := func(length int) bool {
		pre, count := strs[0][:length], len(strs)

		for i := 0; i < count; i++ {
			if pre != strs[i][:length] {
				return false
			}
		}
		return true
	}

	minLen := len(strs[0])
	for _, s := range strs {
		if len(s) < minLen {
			minLen = len(s)
		}
	}

	left, right := 0, minLen

	for left < right {
		mid := (right-left+1)/2 + left

		if isCommonPrefix(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}

	return strs[0][:left]
}

// @lc code=end

// TOOD: 分治法
