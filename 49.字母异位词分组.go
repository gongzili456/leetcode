package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}

	for _, s := range strs {
		stringSlice := strings.Split(s, "")
		sort.Strings(stringSlice)
		sortedKey := strings.Join(stringSlice, "")
		m[sortedKey] = append(m[sortedKey], s)
	}

	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

// @lc code=end
