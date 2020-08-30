package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
func wordPattern(pattern string, str string) bool {
	words := strings.Split(str, " ")

	if len(words) != len(pattern) {
		return false
	}

	// word -> pattern & pattern -> word
	wordMap, PatternMap := map[string]byte{}, map[byte]string{}

	for i := range pattern {
		if word, ok := PatternMap[pattern[i]]; ok {
			if word != words[i] {
				return false
			}
		} else {
			PatternMap[pattern[i]] = words[i]
		}

		if p, ok := wordMap[words[i]]; ok {
			if p != pattern[i] {
				return false
			}
		} else {
			wordMap[words[i]] = pattern[i]
		}
	}

	return true
}

// @lc code=end
