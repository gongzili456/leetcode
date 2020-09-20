package leetcode

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// TODO 单词接龙
// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 预处理，找出单词的通用状态
	dic := map[string][]string{}
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			s := strings.Join(append([]string(nil), w[:i], "*", w[i+1:]), "")
			dic[s] = append(dic[s], w)
		}
	}

	vis := map[string]bool
	q := []string{beginWord}
	depth := 0

	for len(q) > 0 {
		depth++


	}
	return 0
}

// @lc code=end
