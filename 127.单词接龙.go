package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=127 lang=golang
 *
 * [127] 单词接龙
 */

// @lc code=start
func ladderLength(beginWord string, endWord string, wordList []string) int {
	// 预处理，找出单词的通用状态，构建图结构
	dic := map[string][]string{}
	for _, w := range wordList {
		for i := 0; i < len(w); i++ {
			s := strings.Join(append([]string(nil), w[:i], "*", w[i+1:]), "")
			dic[s] = append(dic[s], w)
		}
	}

	type record struct {
		word  string
		depth int
	}

	vis := map[string]bool{}
	q := []record{record{beginWord, 1}}

	// BFS 图遍历
	for len(q) > 0 {
		head := q[0]
		q = q[1:]

		if head.word == endWord {
			return head.depth
		}

		vis[head.word] = true

		for i := 0; i < len(head.word); i++ {
			k := strings.Join(append([]string(nil), head.word[:i], "*", head.word[i+1:]), "")
			if v, ok := dic[k]; ok {
				for _, s := range v {
					if !vis[s] {
						q = append(q, record{word: s, depth: head.depth + 1})
					}
				}
			}
		}
	}

	return 0
}

// @lc code=end
