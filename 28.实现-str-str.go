package leetcode

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	// 双指针
	haystackLen, needleLen := len(haystack), len(needle)

	hp := 0
	// 循环 needleLen 的倍数
	for hp < haystackLen-needleLen+1 {
		// 匹配 needle 的首字符
		for hp < haystackLen-needleLen+1 && haystack[hp] != needle[0] {
			hp++
		}

		currLen, np := 0, 0
		for np < needleLen && hp < haystackLen && haystack[hp] == needle[np] {
			currLen++
			np++
			hp++
		}

		if currLen == needleLen {
			return hp - needleLen
		}

		// 如果匹配失败，退回已经匹配的长度
		hp = hp - currLen + 1
	}

	return -1
}

// @lc code=end
