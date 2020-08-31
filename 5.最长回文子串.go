package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	// 动态规划：一个回文字符串，去除首尾的子串依然是回文
	n := len(s)

	if n < 2 {
		return s
	}

	// 最大长度，起始位置
	max, begin := 1, 0

	// 二维数组矩阵
	dp := make([][]bool, n)

	// 从 0,0 到 i,i 的线值为 true，因为单个字符自己就是一个回文字符串
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// 现在矩阵如下所示

	// j j j j j j
	// T 0 0 0 0 0 i
	// 0 T 0 0 0 0 i
	// 0 0 T 0 0 0 i
	// 0 0 0 T 0 0 i
	// 0 0 0 0 T 0 i
	// 0 0 0 0 0 T
	// 只需要处理右上部分的矩阵即可
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			// 两头不相同，则不是回文
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 { // 长度为 2 时
					dp[i][j] = true
				} else {
					// 两头减 1 的字串是否是回文，重用字串的值
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > max {
				max = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+max]
}

// @lc code=end
